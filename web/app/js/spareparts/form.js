// @ts-check
'use strict';

import { buildJsonData } from '../mappings';
import { isObjectWithEntries } from '../utils';

const form = /** @type {HTMLFormElement} */ (
  document.getElementById('app-form')
);
const methodInput = /** @type {HTMLInputElement} */ (
  form.querySelector('[name="_method"]')
);
const modelInput = /** @type {HTMLInputElement} */ (
  form?.querySelector('[name="_model"]')
);
const alert = /** @type {HTMLElement} */ (document.getElementById('alert'));
const alertMessage = /** @type {HTMLElement} */ (
  document.getElementById('alert-message')
);
const inputElements = /** @type {NodeListOf<HTMLInputElement>} */ (
  document.querySelectorAll('input, select, textarea')
);

inputElements.forEach((element) => {
  element.dataset.originalValue = element.value;
});

form.addEventListener('submit', async function (e) {
  e.preventDefault();

  try {
    let method = this.method;

    if (methodInput) method = methodInput.value;

    method = method.toLocaleUpperCase();

    let payload;
    const formData = new FormData(this);

    if (method === 'PATCH') {
      const changedData = {};

      inputElements.forEach((element) => {
        if (element.value && element.value !== element.dataset.originalValue) {
          changedData[element.name] = element.value;
          element.dataset.originalValue = element.value;
        }
      });

      payload = buildJsonData(changedData, modelInput.value);
    } else {
      payload = buildJsonData(formData, modelInput.value);
    }

    if (!isObjectWithEntries(payload)) return;

    const res = await fetch(this.action, {
      method,
      body: JSON.stringify(payload),
      headers: { 'Content-Type': 'application/json' },
    });

    /** @type {APIResponse} */
    const { success, message } = await res.json();

    if (!success) {
      alertMessage.textContent = message;
      throw new Error(message);
    }

    if (method === 'POST') this.reset();

    alertMessage.textContent = message;
    alert.style.display = 'block';
  } catch (error) {
    console.error('An error occured:', error);
  }
});
