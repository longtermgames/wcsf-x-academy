function escapeHtml(str) {
  var div = document.createElement('div');
  div.textContent = str == null ? '' : String(str);
  return div.innerHTML;
}

fetch('assets/payment.json')
  .then(function (res) { return res.json(); })
  .then(function (payment) {
    var block = document.getElementById('payment-block');
    if (!block) return;
    if (!payment.enabled || !payment.qrImage) return;
    block.innerHTML = (
      '<div class="payment-card reveal">' +
        '<img class="payment-qr" src="assets/images/' + encodeURIComponent(payment.qrImage) + '" alt="QR для оплаты через ' + escapeHtml(payment.provider) + '" loading="lazy" />' +
        '<div class="payment-info">' +
          '<span class="payment-provider">' + escapeHtml(payment.provider) + '</span>' +
          (payment.phone ? '<span class="payment-phone">' + escapeHtml(payment.phone) + '</span>' : '') +
          (payment.note ? '<p class="payment-note">' + escapeHtml(payment.note) + '</p>' : '') +
        '</div>' +
      '</div>'
    );
  })
  .catch(function (err) {
    console.error('payment.json error:', err);
  });
