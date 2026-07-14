document.querySelectorAll('.reveal').forEach(function (el) {
  var io = new IntersectionObserver(function (entries) {
    entries.forEach(function (entry) {
      if (entry.isIntersecting) {
        entry.target.classList.add('in');
        io.unobserve(entry.target);
      }
    });
  }, { threshold: 0.15 });
  io.observe(el);
});

var burger = document.querySelector('.burger');
var nav = document.querySelector('nav.primary');
if (burger) {
  burger.addEventListener('click', function () {
    var open = nav.style.display === 'flex';
    nav.style.cssText = open
      ? ''
      : 'display:flex;flex-direction:column;position:absolute;top:100%;left:0;right:0;background:var(--bg);padding:16px 24px;border-bottom:1px solid var(--border);gap:16px;';
  });
}
