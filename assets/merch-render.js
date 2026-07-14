fetch('assets/products.json')
  .then(function (res) { return res.json(); })
  .then(function (products) {
    var grid = document.getElementById('merch-grid');
    if (!grid) return;
    grid.innerHTML = products.map(function (p) {
      var text = encodeURIComponent(p.message || ('Хочу заказать ' + p.name));
      return (
        '<div class="merch-card">' +
          '<div class="merch-media">' +
            '<img class="tee" src="assets/images/' + p.image + '" alt="' + p.name + '" loading="lazy" />' +
          '</div>' +
          '<div class="merch-body">' +
            '<span class="merch-name">' + p.name + '</span>' +
            '<p class="merch-desc">' + p.desc + '</p>' +
            '<div class="merch-foot">' +
              '<span class="merch-price">' + p.price + ' <small>сом</small></span>' +
              '<a class="btn btn-solid btn-sm" href="https://wa.me/' + p.whatsapp + '?text=' + text + '" target="_blank" rel="noopener">Заказать</a>' +
            '</div>' +
          '</div>' +
        '</div>'
      );
    }).join('');
  })
  .catch(function (err) {
    var grid = document.getElementById('merch-grid');
    if (grid) grid.innerHTML = '<p class="doc-note">Не получилось загрузить список товаров. Проверьте, что файл assets/products.json — валидный JSON.</p>';
    console.error('products.json error:', err);
  });
