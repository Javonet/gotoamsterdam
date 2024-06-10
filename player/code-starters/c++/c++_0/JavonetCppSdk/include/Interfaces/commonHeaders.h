#pragma once
#include <memory>
#include <string>
#include <deque>
#include <any>
#include <vector>

#ifdef _WIN32
#ifdef BUILD_DLL
#define dllExport __declspec(dllexport)
#else
#define dllExport __declspec(dllimport)
#endif
#endif

#if defined(__linux__) || defined(__APPLE__)
#if BUILD_DLL
#define dllExport __attribute__((visibility("default")))
#else
#define dllExport
#endif
#endif