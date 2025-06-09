#ifndef TWO_FER_H_
#define TWO_FER_H_

#include <string>
#include <string_view>

namespace two_fer {

[[nodiscard]]
std::string two_fer(std::string_view name = "you") noexcept;

}  // namespace two_fer

#endif    // TWO_FER_H_