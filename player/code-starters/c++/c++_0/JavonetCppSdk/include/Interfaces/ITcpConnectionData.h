#pragma once
#include "IConnectionData.h"

namespace JavonetNS::Cpp::Utils::Interfaces {
	class ITcpConnectionData : public IConnectionData {
	public:
		virtual ~ITcpConnectionData() = default;
	};
}