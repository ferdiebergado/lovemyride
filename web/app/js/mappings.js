// @ts-check
'use strict';

// Mapping between form fields and model properties (example)
const modelMappings = {
  sparepart: {
    description: 'description',
    maintenance_interval: (/** @type {string} */ value) => Number(value), // convert to number
  },
};

/**
 * Build JSON data based on model.
 *
 * @param {FormData|Object} data
 * @param {string} modelName
 * @returns {Object}
 */
export function buildJsonData(data, modelName) {
  const mapping = modelMappings[modelName];

  const jsonData = {};

  for (const key in mapping) {
    let rawValue;

    if (data instanceof FormData) {
      rawValue = data.get(key);
    } else {
      rawValue = data[key];
    }

    if (rawValue) {
      if (typeof mapping[key] === 'function') {
        jsonData[key] = mapping[key](rawValue);
      } else {
        jsonData[key] = rawValue;
      }
    }
  }

  return jsonData;
}
