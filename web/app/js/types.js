// @ts-check
'use strict';

/**
 * @typedef {Object} ValidationError
 *
 * @property {string} field
 * @property {string} error
 */

/**
 * @typedef {Object} APIResponse
 *
 * @property {boolean} success
 * @property {string} message
 * @property {ValidationError[]} errors
 * @property {Object} data
 * @property {Object} meta
 */

/**
 * @typedef {Object} SparePartRequest
 *
 * @property {string} description
 * @property {number} maintenanceInterval
 */
