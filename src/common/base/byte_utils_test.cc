#include <gmock/gmock.h>
#include <gtest/gtest.h>

#include <random>

#include "src/common/base/byte_utils.h"
#include "src/common/base/types.h"

namespace pl {
namespace utils {

TEST(UtilsTest, TestReverseArrayBytes) {
  {
    char result[4];
    char input[] = {'\x12', '\x34', '\x56', '\x78'};
    ReverseBytes(input, result);
    char expected[] = "\x78\x56\x34\x12";
    EXPECT_EQ(result[0], expected[0]);
    EXPECT_EQ(result[1], expected[1]);
    EXPECT_EQ(result[2], expected[2]);
    EXPECT_EQ(result[3], expected[3]);
  }

  {
    uint8_t result[4];
    uint8_t input[] = {'\x12', '\x34', '\x56', '\x78'};
    ReverseBytes(input, result);
    uint8_t expected[] = "\x78\x56\x34\x12";
    EXPECT_EQ(result[0], expected[0]);
    EXPECT_EQ(result[1], expected[1]);
    EXPECT_EQ(result[2], expected[2]);
    EXPECT_EQ(result[3], expected[3]);
  }
}

TEST(UtilsTest, TestReverseTypeBytes) {
  {
    int32_t x = 0x01020304;
    int32_t y = ReverseBytes(x);
    EXPECT_EQ(y, 0x04030201);
  }
  {
    int32_t x = 0x01020384;
    int32_t y = ReverseBytes(x);
    EXPECT_EQ(y, 0x84030201);
    EXPECT_EQ(y, -0x7bfcfdff);
  }
  {
    int64_t x = 0x01020384;
    int64_t y = ReverseBytes(x);
    EXPECT_EQ(y, 0x8403020100000000);
    EXPECT_EQ(y, -0x7bfcfdff00000000);
  }
}

TEST(UtilsTest, TestReverseBytesInvertability) {
  std::default_random_engine rng;
  std::uniform_int_distribution<uint32_t> uniform_uint32_dist;
  std::uniform_real_distribution<double> uniform_double_dist;

  for (int i = 0; i < 100000; ++i) {
    {
      int32_t x = uniform_uint32_dist(rng);
      int32_t y = ReverseBytes(x);
      EXPECT_EQ(ReverseBytes(y), x);
    }

    {
      double x = uniform_double_dist(rng);
      double y = ReverseBytes(x);
      EXPECT_EQ(ReverseBytes(y), x);
    }
  }
}

TEST(UtilsTest, TestLEndianBytesToInt) {
  // uint32_t cases.
  EXPECT_EQ(LEndianBytesToInt<uint32_t>(ConstString("\x78\x56\x34\x12")), 0x12345678);
  EXPECT_EQ((LEndianBytesToInt<uint32_t, 3>(ConstString("\xc6\x00\x00"))), 0x0000c6);
  EXPECT_EQ(LEndianBytesToInt<uint32_t>(ConstString("\x33\x77\xbb\xff")), 0xffbb7733);

  // int32_t cases.
  EXPECT_EQ(LEndianBytesToInt<int32_t>(ConstString("\x78\x56\x34\x12")), 0x12345678);
  EXPECT_EQ((LEndianBytesToInt<int32_t, 3>(ConstString("\xc6\x00\x00"))), 0x0000c6);
  EXPECT_EQ(LEndianBytesToInt<int32_t>(ConstString("\x33\x77\xbb\xff")), -0x4488cd);

  // 64-bit cases.
  EXPECT_EQ(
      LEndianBytesToInt<int64_t>(std::string(ConstStringView("\xf0\xde\xbc\x9a\x78\x56\x34\x12"))),
      0x123456789abcdef0);
  EXPECT_EQ(
      LEndianBytesToInt<int64_t>(std::string(ConstStringView("\xf0\xde\xbc\x9a\x78\x56\x34\xf2"))),
      -0xdcba98765432110);
}

TEST(UtilsTest, TestLEndianBytesToFloat) {
  std::string float_bytes = ConstString("\x33\x33\x23\x41");
  std::string double_bytes = ConstString("\x66\x66\x66\x66\x66\x66\x24\x40");

  EXPECT_FLOAT_EQ(LEndianBytesToFloat<float>(float_bytes), 10.2f);
  EXPECT_DOUBLE_EQ(LEndianBytesToFloat<double>(double_bytes), 10.2);
}

TEST(UtilsTest, TestLEndianBytesToFloatUnAligned) {
  std::string float_bytes = ConstString("\x33\x33\x23\x41");
  std::string double_bytes = ConstString("\x66\x66\x66\x66\x66\x66\x24\x40");

  std::string unaligned_float_bytes = ConstString("\x00") + float_bytes;
  std::string unaligned_double_bytes = ConstString("\x00") + double_bytes;

  EXPECT_FLOAT_EQ(LEndianBytesToFloat<float>(unaligned_float_bytes.substr(1)), 10.2f);
  EXPECT_DOUBLE_EQ(LEndianBytesToFloat<double>(unaligned_double_bytes.substr(1)), 10.2);
}

TEST(UtilsTest, TestIntToLEndianBytes) {
  {
    char result[4];
    IntToLEndianBytes(0x12345678, result);
    char expected[] = "\x78\x56\x34\x12";
    EXPECT_EQ(result[0], expected[0]);
    EXPECT_EQ(result[1], expected[1]);
    EXPECT_EQ(result[2], expected[2]);
    EXPECT_EQ(result[3], expected[3]);
  }

  {
    char result[3];
    IntToLEndianBytes(198, result);
    char expected[] = "\xc6\x00\x00";
    EXPECT_EQ(result[0], expected[0]);
    EXPECT_EQ(result[1], expected[1]);
    EXPECT_EQ(result[2], expected[2]);
  }

  {
    uint8_t result[4];
    IntToLEndianBytes(0x12345678, result);
    uint8_t expected[] = "\x78\x56\x34\x12";
    EXPECT_EQ(result[0], expected[0]);
    EXPECT_EQ(result[1], expected[1]);
    EXPECT_EQ(result[2], expected[2]);
    EXPECT_EQ(result[3], expected[3]);
  }

  {
    uint8_t result[8];
    IntToLEndianBytes(0x123456789abcdef0, result);
    uint8_t expected[] = "\xf0\xde\xbc\x9a\x78\x56\x34\x12";
    EXPECT_EQ(result[0], expected[0]);
    EXPECT_EQ(result[1], expected[1]);
    EXPECT_EQ(result[2], expected[2]);
    EXPECT_EQ(result[3], expected[3]);
    EXPECT_EQ(result[4], expected[4]);
    EXPECT_EQ(result[5], expected[5]);
    EXPECT_EQ(result[6], expected[6]);
    EXPECT_EQ(result[7], expected[7]);
  }
}

TEST(UtilsTest, TestBEndianBytesToInt) {
  // uint32_t cases.
  EXPECT_EQ(BEndianBytesToInt<uint32_t>(ConstString("\x12\x34\x56\x78")), 0x12345678);
  EXPECT_EQ((BEndianBytesToInt<uint32_t, 3>(ConstString("\x00\x00\xc6"))), 0x0000c6);
  EXPECT_EQ(BEndianBytesToInt<uint32_t>(ConstString("\xff\xbb\x77\x33")), 0xffbb7733);

  // int32_t cases.
  EXPECT_EQ(BEndianBytesToInt<int32_t>(ConstString("\x12\x34\x56\x78")), 0x12345678);
  EXPECT_EQ((BEndianBytesToInt<int32_t, 3>(ConstString("\x00\x00\xc6"))), 0x0000c6);
  EXPECT_EQ(BEndianBytesToInt<int32_t>(ConstString("\xff\xbb\x77\x33")), -0x4488cd);

  // 64-bit cases.
  EXPECT_EQ(
      BEndianBytesToInt<int64_t>(std::string(ConstStringView("\x12\x34\x56\x78\x9a\xbc\xde\xf0"))),
      0x123456789abcdef0);
  EXPECT_EQ(
      BEndianBytesToInt<int64_t>(std::string(ConstStringView("\xf2\x34\x56\x78\x9a\xbc\xde\xf0"))),
      -0xdcba98765432110);
}

TEST(UtilsTest, TestBEndianBytesToFloat) {
  std::string float_bytes = ConstString("\x41\x23\x33\x33");
  std::string double_bytes = ConstString("\x40\x24\x66\x66\x66\x66\x66\x66");

  EXPECT_FLOAT_EQ(BEndianBytesToFloat<float>(float_bytes), 10.2f);
  EXPECT_DOUBLE_EQ(BEndianBytesToFloat<double>(double_bytes), 10.2);
}

TEST(UtilsTest, TestBEndianBytesToFloatUnAligned) {
  std::string float_bytes = ConstString("\x41\x23\x33\x33");
  std::string double_bytes = ConstString("\x40\x24\x66\x66\x66\x66\x66\x66");

  std::string unaligned_float_bytes = ConstString("\x00") + float_bytes;
  std::string unaligned_double_bytes = ConstString("\x00") + double_bytes;

  EXPECT_FLOAT_EQ(BEndianBytesToFloat<float>(unaligned_float_bytes.substr(1)), 10.2f);
  EXPECT_DOUBLE_EQ(BEndianBytesToFloat<double>(unaligned_double_bytes.substr(1)), 10.2);
}

}  // namespace utils
}  // namespace pl
