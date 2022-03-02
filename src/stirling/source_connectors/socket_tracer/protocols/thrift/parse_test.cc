/*
 * Copyright 2018- The Pixie Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

#include <string>
#include <utility>
#include "src/stirling/source_connectors/socket_tracer/protocols/thrift/parse.h"

#include "src/common/testing/testing.h"

namespace px {
namespace stirling {
namespace protocols { 
    constexpr uint8_t thriftData[] = {
    // Thrift data
    0x80, 0x01, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x67, 0x65, 0x74, 0x00, 0x00,
    0x00, 0x00, 0x0a, 0x00, 0x00, 0x14, 0x00, 0x35, 0xc9, 0xd0, 0x80, 0x00, 0x00, 0x00,
    };

    class ThriftParserTest : public ::testing::Test {};

    // TEST(ThriftParserTest, SimpleTest) {
    //     ASSERT_EQ(thrift::HelloWorld(),5);
    // }

    TEST_F(ThriftParserTest, ParseSimpleFrame) {

        auto frame_view = CreateStringView<char>(CharArrayStringView<uint8_t>(thriftData));
        ParseState state = thrift::ParseFrame(&frame_view);
        ASSERT_EQ(state, ParseState::kSuccess); 
    }
    
}  // namespace protocols
}  // namespace stirling
}
