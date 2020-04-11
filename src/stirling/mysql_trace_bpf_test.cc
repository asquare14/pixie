#include "src/stirling/mysql/test_data.h"
#include "src/stirling/mysql/test_utils.h"
#include "src/stirling/testing/common.h"
#include "src/stirling/testing/socket_trace_bpf_test_fixture.h"

namespace pl {
namespace stirling {

using ::pl::stirling::testing::FindRecordIdxMatchesPid;
using ::pl::stirling::testing::SocketTraceBPFTest;
using ::pl::stirling::testing::TCPSocket;
using ::pl::types::ColumnWrapper;
using ::testing::IsEmpty;
using ::testing::SizeIs;

testing::SendRecvScript GetPrepareExecuteScript() {
  std::vector<std::string_view> prepare_req_packets(mysql::testdata::kRawStmtPrepareReq.begin(),
                                                    mysql::testdata::kRawStmtPrepareReq.end());
  std::vector<std::string_view> prepare_resp_packets(mysql::testdata::kRawStmtPrepareResp.begin(),
                                                     mysql::testdata::kRawStmtPrepareResp.end());
  std::vector<std::string_view> execute_req_packets(mysql::testdata::kRawStmtExecuteReq.begin(),
                                                    mysql::testdata::kRawStmtExecuteReq.end());
  std::vector<std::string_view> execute_resp_packets(mysql::testdata::kRawStmtExecuteResp.begin(),
                                                     mysql::testdata::kRawStmtExecuteResp.end());
  std::vector<std::string_view> close_req_packets(mysql::testdata::kRawStmtCloseReq.begin(),
                                                  mysql::testdata::kRawStmtCloseReq.end());

  testing::SendRecvScript script = {{prepare_req_packets, prepare_resp_packets},
                                    {execute_req_packets, execute_resp_packets},
                                    {close_req_packets, {""}}};

  return script;
}

testing::SendRecvScript GetQueryScript() {
  std::vector<std::string_view> query_req_packets(mysql::testdata::kRawQueryReq.begin(),
                                                  mysql::testdata::kRawQueryReq.end());
  std::vector<std::string_view> query_resp_packets(mysql::testdata::kRawQueryResp.begin(),
                                                   mysql::testdata::kRawQueryResp.end());

  testing::SendRecvScript script = {{query_req_packets, query_resp_packets}};
  return script;
}

TEST_F(SocketTraceBPFTest, MySQLStmtPrepareExecuteClose) {
  ConfigureCapture(TrafficProtocol::kProtocolMySQL, kRoleClient);
  testing::ClientServerSystem system;
  system.RunClientServer<&TCPSocket::Read, &TCPSocket::Write>(GetPrepareExecuteScript());

  // Check that HTTP table did not capture any data.
  {
    DataTable data_table(kHTTPTable);
    source_->TransferData(ctx_.get(), kHTTPTableNum, &data_table);
    types::ColumnWrapperRecordBatch& record_batch = *data_table.ActiveRecordBatch();

    EXPECT_THAT(FindRecordIdxMatchesPid(record_batch, kMySQLUPIDIdx, getpid()), IsEmpty());
  }

  // Check that MySQL table did capture the appropriate data.
  {
    DataTable data_table(kMySQLTable);
    source_->TransferData(ctx_.get(), kMySQLTableNum, &data_table);
    types::ColumnWrapperRecordBatch& record_batch = *data_table.ActiveRecordBatch();

    const std::vector<size_t> target_record_indices =
        FindRecordIdxMatchesPid(record_batch, kMySQLUPIDIdx, getpid());
    ASSERT_THAT(target_record_indices, SizeIs(3));

    EXPECT_EQ(
        "SELECT sock.sock_id AS id, GROUP_CONCAT(tag.name) AS tag_name FROM "
        "sock "
        "JOIN sock_tag ON "
        "sock.sock_id=sock_tag.sock_id JOIN tag ON sock_tag.tag_id=tag.tag_id WHERE tag.name=? "
        "GROUP "
        "BY id ORDER BY ?",
        record_batch[kMySQLReqBodyIdx]->Get<types::StringValue>(target_record_indices[0]));

    EXPECT_EQ(
        "SELECT sock.sock_id AS id, GROUP_CONCAT(tag.name) AS tag_name FROM "
        "sock "
        "JOIN sock_tag ON "
        "sock.sock_id=sock_tag.sock_id JOIN tag ON sock_tag.tag_id=tag.tag_id WHERE tag.name=brown "
        "GROUP "
        "BY id ORDER BY id",
        record_batch[kMySQLReqBodyIdx]->Get<types::StringValue>(target_record_indices[1]));

    EXPECT_EQ("",
              record_batch[kMySQLReqBodyIdx]->Get<types::StringValue>(target_record_indices[2]));
  }
}

TEST_F(SocketTraceBPFTest, MySQLQuery) {
  ConfigureCapture(TrafficProtocol::kProtocolMySQL, kRoleClient);
  testing::ClientServerSystem system;
  system.RunClientServer<&TCPSocket::Read, &TCPSocket::Write>(GetQueryScript());

  // Check that HTTP table did not capture any data.
  {
    DataTable data_table(kHTTPTable);
    source_->TransferData(ctx_.get(), kHTTPTableNum, &data_table);
    types::ColumnWrapperRecordBatch& record_batch = *data_table.ActiveRecordBatch();

    EXPECT_THAT(FindRecordIdxMatchesPid(record_batch, kMySQLUPIDIdx, getpid()), IsEmpty());
  }

  // Check that MySQL table did capture the appropriate data.
  {
    DataTable data_table(kMySQLTable);
    source_->TransferData(ctx_.get(), kMySQLTableNum, &data_table);
    types::ColumnWrapperRecordBatch& record_batch = *data_table.ActiveRecordBatch();

    const std::vector<size_t> target_record_indices =
        FindRecordIdxMatchesPid(record_batch, kMySQLUPIDIdx, getpid());
    ASSERT_THAT(target_record_indices, SizeIs(1));

    EXPECT_EQ("SELECT name FROM tag;",
              record_batch[kMySQLReqBodyIdx]->Get<types::StringValue>(target_record_indices[0]));
  }
}

}  // namespace stirling
}  // namespace pl
