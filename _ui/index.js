import "./index.css"
import TomSelect from '@tabler/core/dist/libs/tom-select/dist/js/tom-select.base.min.js'

document.addEventListener("DOMContentLoaded", function () {
  (new TomSelect('#select-users', {
    copyClassesToDropdown: false,
    dropdownClass: 'dropdown-menu ts-dropdown',
    optionClass: 'dropdown-item',
    controlInput: '<input>',
    render: {
      item: function (data, escape) {
        if (data.customProperties) {
          return '<div><span class="dropdown-item-indicator">' + data.customProperties + '</span>' + escape(data.text) + '</div>';
        }
        return '<div>' + escape(data.text) + '</div>';
      },
      option: function (data, escape) {
        if (data.customProperties) {
          return '<div><span class="dropdown-item-indicator">' + data.customProperties + '</span>' + escape(data.text) + '</div>';
        }
        return '<div>' + escape(data.text) + '</div>';
      },
    },
  }));

  (new TomSelect('#select-remote', {
    copyClassesToDropdown: false,
    dropdownClass: 'dropdown-menu ts-dropdown',
    optionClass: 'dropdown-item',
    controlInput: '<input>',
    valueField: 'url',
    labelField: 'name',
    searchField: 'name',
    load: function (query, callback) {
      var url = 'https://api.github.com/search/repositories?q=' + encodeURIComponent(query);
      fetch(url)
        .then(response => response.json())
        .then(json => {
          callback(json.items);
        }).catch(() => {
          callback();
        });
    },
    render: {
      option: function (item, escape) {
        return `<div class="py-2 d-flex">
                <div class="icon me-3">
                  <img class="img-fluid" src="${item.owner.avatar_url}" />
                </div>
                <div>
                  <div class="mb-1">
                    <span class="h4">
                      ${escape(item.name)}
                    </span>
                    <span class="text-muted">by ${escape(item.owner.login)}</span>
                  </div>
                   <div class="description">${escape(item.description)}</div>
                </div>
              </div>`;
      },
      item: function (item, escape) {
        return `<div class="py-2 d-flex">
                <div class="icon me-3">
                  <img class="img-fluid" src="${item.owner.avatar_url}" />
                </div>
                <div>
                  <div class="mb-1">
                    <span class="h4">
                      ${escape(item.name)}
                    </span>
                    <span class="text-muted">by ${escape(item.owner.login)}</span>
                  </div>
                   <div class="description">${escape(item.description)}</div>
                </div>
              </div>`;
      }
    },
  }));

});