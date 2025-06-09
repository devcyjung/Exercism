const RESPONSES = {
  Question: "Sure.",
  Yelling: "Whoa, chill out!",
  YellingQuestion: "Calm down, I know what I'm doing!",
  Silence: "Fine. Be that way!",
  Otherwise: "Whatever.",
} as const

type Response = typeof RESPONSES[keyof typeof RESPONSES]

export function hey(message: string): Response {
  const trimmed = message.trim()
  if (trimmed.length === 0) {
    return RESPONSES.Silence
  }
  const isAsking = trimmed.endsWith("?")
  const upper = trimmed.toUpperCase()
  const isYelling = trimmed.toLowerCase() !== upper && trimmed === upper
  switch (true) {
    case isYelling && isAsking:
      return RESPONSES.YellingQuestion
    case isYelling:
      return RESPONSES.Yelling
    case isAsking:
      return RESPONSES.Question
    default:
      return RESPONSES.Otherwise
  }
}