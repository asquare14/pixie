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
    
        PL_ASSIGN_OR(tm.version, decoder.ExtractInt<uint16_t>(),
                 return ParseState::kInvalid);
       
        if(tm.version != 0x8001) {
            std::cout<<"Payload does not conform to thrift standard:"<< tm.version<<"\n";
        }


        PL_ASSIGN_OR(tm.unused, decoder.ExtractInt<uint8_t>(),
                 return ParseState::kInvalid);
        // std::cout<<tm.unused<<" :unused\n";

         PL_ASSIGN_OR(tm.message_type, decoder.ExtractInt<uint8_t>(),
                 return ParseState::kInvalid);        
        // std::cout<<tm.message_type<<" :message_type\n";
        if( tm.message_type <1 || tm.message_type>4) {
            std::cout<<"Payload does not conform to thrift standard:"<< tm.message_type<<" leading bits in message type is less than 5\n";
        }

        PL_ASSIGN_OR(tm.name_length, decoder.ExtractInt<int32_t>(),
                 return ParseState::kInvalid);
        //  std::cout<<"name len: "<<tm.name_length<<"\n"; 
         if( tm.name_length<0) {
            std::cout<<"Payload does not conform to thrift standard:"<< tm.name_length<<"\n";
        }

        PL_ASSIGN_OR(tm.name, decoder.ExtractInt<int64_t>(),
                 return ParseState::kInvalid);
        // std::cout<<"name: "<<tm.name<<"\n"; 


        PL_ASSIGN_OR(tm.sequence_id, decoder.ExtractInt<int32_t>(),
                 return ParseState::kInvalid);
        // std::cout<<"seq_id: "<<tm.sequence_id<<"\n";    

        while(1) {
            int8_t  fieldType = 0;
            PL_ASSIGN_OR(fieldType, decoder.ExtractInt<int8_t>(),
                 return ParseState::kInvalid);
            std::cout<<fieldType<<"fieldType\n";
            
            if(fieldType==0) {
                std::cout<<"Break\n";
                 break;
            }

            int16_t fieldID = 0;
             PL_ASSIGN_OR(fieldID, decoder.ExtractInt<int16_t>(),
                 return ParseState::kInvalid);
            std::cout<<fieldID<<"fieldID\n";

            switch(fieldType) {
                        case ThirftBool: {
                //             bool booleanValue = false;
                //              PL_ASSIGN_OR(booleanValue, decoder.ExtractInt<bool>(),
                //  return ParseState::kInvalid);
                // hi
                            continue;
                        }
                        case ThriftInt8: {
                            continue;

                        }
                        case ThriftDouble: {
                            continue;

                        }
                        // case ThriftInt16: {
                        //     continue;

                        // }
                        case ThriftInt32: {
                            continue;

                        }
                        // case ThriftInt64: {
                        //     continue;

                        // }
                        case ThriftStringAndBinary: {
                            continue;

                        }
                        case ThriftStruct: {
                            continue;

                        }
                        case ThriftMap: {
                            continue;

                        }
                        case ThriftSet: {
                            continue;

                        } 
                        case ThriftList: {
                            continue;

                        }
                        default: {
                            std::cout<<"None of the types match\n\n";
                        }
            }
            
        }



        return ParseState::kSuccess;
    }

}
}
}
}