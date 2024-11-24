// @ts-check
'use strict';

/**
 * @typedef {Object} TableHeader
 *
 * @property {string} label
 * @property {string} field
 */

const apiPrefix = '/api';
const datatable = /** @type {HTMLTableElement} */ (
  document.getElementById('datatable')
);
const tableHead = datatable.querySelector('thead');
const tableBody = datatable.querySelector('tbody');

const endpoint = datatable.dataset.url;
const headers = /** @type {TableHeader[]} */ (
  JSON.parse(datatable?.dataset.headers || '[]')
);
const data = /** @type {Object[]} */ (
  JSON.parse(datatable?.dataset.data || '[]')
);

function renderTableHead() {
  const theadFragment = document.createDocumentFragment();
  const row = document.createElement('tr');

  headers.forEach((header) => {
    const th = document.createElement('th');
    th.textContent = header.label;
    row.appendChild(th);
  });

  const th = document.createElement('th');
  th.textContent = 'Actions';
  row.appendChild(th);

  theadFragment.appendChild(row);

  tableHead?.appendChild(theadFragment);
}

// Render table with optimized DOM manipulation
function renderTableBody() {
  const bodyFragment = document.createDocumentFragment();

  if (data.length > 0) {
    data.forEach((row) => {
      const tr = document.createElement('tr');

      headers.forEach(({ field }) => {
        const td = document.createElement('td');
        td.textContent = row[field];

        tr.appendChild(td);
      });

      const td = document.createElement('td');

      const viewLink = document.createElement('a');
      viewLink.classList.add('btn', 'btn-primary');
      viewLink.href = `${endpoint?.replace(apiPrefix, '')}/${row.id}`;
      viewLink.textContent = 'Info';

      const editLink = document.createElement('a');
      editLink.classList.add('btn', 'btn-secondary');
      editLink.href = `${endpoint?.replace(apiPrefix, '')}/${row.id}/edit`;
      editLink.textContent = 'Edit';

      td.appendChild(viewLink);
      td.appendChild(editLink);
      tr.appendChild(td);

      bodyFragment.appendChild(tr);
    });
  } else {
    const noDataRow = document.createElement('tr');
    const noDataCell = document.createElement('td');

    noDataCell.colSpan = headers.length + 1;
    noDataCell.textContent = 'No data.';
    noDataRow.appendChild(noDataCell);
    bodyFragment.appendChild(noDataRow);
  }

  if (tableBody) {
    tableBody.innerHTML = '';
    tableBody.appendChild(bodyFragment);
  }
}

renderTableHead();
renderTableBody();
