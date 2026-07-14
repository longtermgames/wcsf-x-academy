fetch('assets/documents.json')
  .then(function (res) { return res.json(); })
  .then(function (groups) {
    var root = document.getElementById('doc-groups');
    if (!root) return;
    root.innerHTML = groups.map(function (group) {
      var rows = group.items.map(function (doc) {
        var hasFile = !!doc.file;
        var href = hasFile ? 'assets/docs/' + doc.file : '#';
        var meta = hasFile ? (doc.updated ? 'Обновлено ' + doc.updated : '') + (doc.size ? ' · ' + doc.size : '') : '—';
        return (
          '<a class="doc-row" href="' + href + '"' + (hasFile ? ' download' : '') + '>' +
            '<span class="doc-icon">PDF</span>' +
            '<span>' +
              '<span class="doc-title">' + doc.title + '</span>' +
              '<span class="doc-desc">' + doc.desc + '</span>' +
            '</span>' +
            '<span class="doc-meta">' + meta + '</span>' +
            '<span class="doc-dl">↓</span>' +
          '</a>'
        );
      }).join('');
      return (
        '<h2 class="doc-group-label">' + group.group + '</h2>' +
        '<div class="doc-board reveal">' + rows + '</div>'
      );
    }).join('');
  })
  .catch(function (err) {
    var root = document.getElementById('doc-groups');
    if (root) root.innerHTML = '<p class="doc-note">Не получилось загрузить список документов. Проверьте, что файл assets/documents.json — валидный JSON.</p>';
    console.error('documents.json error:', err);
  });
