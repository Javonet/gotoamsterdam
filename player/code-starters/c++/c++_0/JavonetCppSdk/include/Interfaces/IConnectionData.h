#pragma once

#ifndef byte
typedef unsigned char byte;
#endif

#include <string>
#include <vector>

namespace JavonetNS::Cpp::Utils::Interfaces {
	class IConnectionData {
	public:
		virtual ~IConnectionData() = default;
		virtual std::string getFullAddress() = 0;
		virtual std::vector<byte> getAddressBytes() = 0;
		virtual std::vector<byte> getPortBytes() = 0;
	};
}