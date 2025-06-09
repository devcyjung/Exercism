// @ts-check

/**
 * Given a certain command, help the chatbot recognize whether the command is valid or not.
 *
 * @param {string} command
 * @returns {boolean} whether or not is the command valid
 */

const validPattern = /^chatbot/gi

export function isValidCommand(command) {
  const result = validPattern.test(command)
  validPattern.lastIndex = 0
  return result
}

/**
 * Given a certain message, help the chatbot get rid of all the emoji's encryption through the message.
 *
 * @param {string} message
 * @returns {string} The message without the emojis encryption
 */

const emojiPattern = new RegExp(/emoji\d+/, 'gi')

export function removeEmoji(message) {
  return message.replace(emojiPattern, '')
}

/**
 * Given a certain phone number, help the chatbot recognize whether it is in the correct format.
 *
 * @param {string} number
 * @returns {string} the Chatbot response to the phone Validation
 */

const phoneNumberPattern = /^\(\+\d{2}\) \d{3}-\d{3}-\d{3}$/
const successMessage = 'Thanks! You can now download me to your phone.'
const failureMessage = number => `Oops, it seems like I can't reach out to ${number}`

export function checkPhoneNumber(number) {
  return phoneNumberPattern.test(number) ? successMessage : failureMessage(number);
}

/**
 * Given a certain response from the user, help the chatbot get only the URL.
 *
 * @param {string} userInput
 * @returns {string[] | null} all the possible URL's that the user may have answered
 */

const urlPattern = /\w+\.\w+/gi

export function getURL(userInput) {
  return userInput.match(urlPattern)
}

/**
 * Greet the user using the full name data from the profile.
 *
 * @param {string} fullName
 * @returns {string} Greeting from the chatbot
 */

const namePattern = /^(?<surName>\w+), (?<firstName>\w+)$/
const greetingMessage = (firstName, lastName) => `Nice to meet you, ${firstName} ${lastName}`

export function niceToMeetYou(fullName) {
  const {firstName, surName} = fullName.match(namePattern).groups
  return greetingMessage(firstName, surName)
}
