/// <reference path="./global.d.ts" />
// @ts-check

export class TranslationService {
  /**
   * Creates a new service
   * @param {ExternalApi} api the original api
   */
  constructor(api) {
    this.api = api
  }

  /**
   * Attempts to retrieve the translation for the given text.
   *
   * - Returns whichever translation can be retrieved, regardless the quality
   * - Forwards any error from the translation api
   *
   * @param {string} text
   * @returns {Promise<string>}
   */
  async free(text) {
    const response = await this.api.fetch(text)
    return response.translation
  }

  /**
   * Batch translates the given texts using the free service.
   *
   * - Resolves all the translations (in the same order), if they all succeed
   * - Rejects with the first error that is encountered
   * - Rejects with a BatchIsEmpty error if no texts are given
   *
   * @param {string[]} texts
   * @returns {Promise<string[]>}
   */
  async batch(texts) {
    if (texts.length === 0) {
      throw new BatchIsEmpty()
    }
    const responses = await Promise.all(texts.map(text => this.api.fetch(text)))
    return responses.map(resp => resp.translation)
  }

  /**
   * Requests the service for some text to be translated.
   *
   * Note: the request service is flaky, and it may take up to three times for
   *       it to accept the request.
   *
   * @param {string} text
   * @returns {Promise<void>}
   */
  async request(text) {
    for (let retry = 0; retry < 3; ++retry) {
      const {promise, reject, resolve} = Promise.withResolvers()
      this.api.request(text, (result) => {
        switch (result) {
          case undefined:
            resolve()
            break
          default:
            reject(result)
        }
      })
      try {
        return await promise
      } catch (error) {
        if (retry === 2) {
          return Promise.reject(error)
        }
      }
    }
  }

  /**
   * Retrieves the translation for the given text
   *
   * - Rejects with an error if the quality can not be met
   * - Requests a translation if the translation is not available, then retries
   *
   * @param {string} text
   * @param {number} minimumQuality
   * @returns {Promise<string>}
   */
  async premium(text, minimumQuality) {
    const resolveAndInspect = async () => {
      const response = await this.api.fetch(text)
      if (response.quality >= minimumQuality) {
        return response.translation
      }
      return Promise.reject(new QualityThresholdNotMet(text))
    }
    try {
      return await resolveAndInspect()
    } catch (error) {
      if (!(error instanceof QualityThresholdNotMet)) {
        await this.request(text)
        return await resolveAndInspect()
      }
      return Promise.reject(error)
    }
  }
}

/**
 * This error is used to indicate a translation was found, but its quality does
 * not meet a certain threshold. Do not change the name of this error.
 */
export class QualityThresholdNotMet extends Error {
  /**
   * @param {string} text
   */
  constructor(text) {
    super(
      `
The translation of ${text} does not meet the requested quality threshold.
    `.trim(),
    )

    this.text = text
  }
}

/**
 * This error is used to indicate the batch service was called without any
 * texts to translate (it was empty). Do not change the name of this error.
 */
export class BatchIsEmpty extends Error {
  constructor() {
    super(
      `
Requested a batch translation, but there are no texts in the batch.
    `.trim(),
    )
  }
}
