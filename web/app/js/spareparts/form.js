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
const submitBtn = /** @type {HTMLButtonElement} */ (
  form.querySelector('[type="submit"]')
);

let loading = false;

updateInputOriginalValues();

form.removeEventListener('submit', handleFormSubmit);
form.addEventListener('submit', handleFormSubmit);

/**
 * Handle the form submission.
 *
 * @param {SubmitEvent} e
 * @returns
 */
async function handleFormSubmit(e) {
  e.preventDefault();

  if (!form || !submitBtn || !alertMessage || !alert) return;

  loading = true;

  updateSubmitBtn();

  try {
    const method = (methodInput.value || this.method).toLocaleUpperCase();
    const formData = new FormData(this);

    let payload;

    if (method === 'PATCH') {
      payload = buildChangedPayload();
    } else {
      payload = buildJsonData(formData, modelInput?.value);
    }

    if (!isObjectWithEntries(payload)) return;

    const res = await fetch(this.action, {
      method,
      body: JSON.stringify(payload),
      headers: { 'Content-Type': 'application/json' },
    });

    if (!res.ok) {
      throw new Error(`HTTP error! status: ${res.status}`);
    }

    /** @type {import('../types').APIResponse} */
    const { success, message } = await res.json();

    if (!success) {
      alertMessage.textContent = message;
      throw new Error(message);
    }

    if (method === 'POST') {
      this.reset();

      updateInputOriginalValues();
    }

    alertMessage.textContent = message;
    alert.style.display = 'block';
  } catch (error) {
    console.error('An error occured:', error);
    showError(error.message);
  } finally {
    loading = false;
    updateSubmitBtn();
  }
}

function updateSubmitBtn() {
  if (loading) {
    submitBtn.textContent = '';
    submitBtn.disabled = true;
    return;
  }

  submitBtn.textContent = 'Submit';
  submitBtn.disabled = false;
}

/**
 * Display an alert containing an error message.
 *
 * @param {string} msg
 */
function showError(msg) {
  alertMessage.textContent = msg;
  alert.style.display = 'block';
}

function buildChangedPayload() {
  const changedData = {};

  inputElements.forEach((element) => {
    if (element.value && element.value !== element.dataset.originalValue) {
      changedData[element.name] = element.value;
      element.dataset.originalValue = element.value;
    }
  });

  return buildJsonData(changedData, modelInput?.value);
}

function updateInputOriginalValues() {
  inputElements.forEach((element) => {
    element.dataset.originalValue = element.value;
  });
}
