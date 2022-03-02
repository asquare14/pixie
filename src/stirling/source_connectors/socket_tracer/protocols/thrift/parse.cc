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
#include "src/stirling/source_connectors/socket_tracer/protocols/thrift/parse.h"

#include <string>
#include <utility>

namespace px {
namespace stirling {
namespace protocols {
namespace thrift {

    struct ThriftMessage {
        uint16_t version;
        uint8_t  unused;
        uint8_t message_type;
        int32_t name_length;
        int64_t name;
        int32_t sequence_id;
    };
    
    enum {
        ThirftBool =  2,
        ThriftInt8 = 3,
        ThriftDouble = 4,
        ThirftInt16 = 6,
        ThriftInt32 = 8,
        ThirftInt64 = 10,
        ThriftStringAndBinary = 11,
        ThriftStruct = 12,
        ThriftMap = 13,
        ThriftSet = 14,
        ThriftList = 15
    };

    int HelloWorld() {
        return 5;
    }

    ParseState ParseFrame(std::string_view* buf) {
        
       struct ThriftMessage tm;
       
        if(buf->size()<12) {
            std::cout<<"Payload too small to contain thrift data\n";
            return ParseState::kInvalid;
        } 

        // std::cout<<buf->size()<<" buf size:\n";
        
        BinaryDecoder decoder(*buf);

        tm.version = decoder.ExtractInt<uint16_t>().ValueOr(0xffff);

        // std::cout<<tm.version <<"tm.version \n\n";
       
        if(tm.version != 0x8001) {
            std::cout<<"Payload does not conform to thrift standard:"<< tm.version<<"\n";
        }

       
        return ParseState::kSuccess;
    }

}
}
}
}