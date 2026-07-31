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
var header = document.querySelector('header.site');
var navLinks = document.querySelectorAll('.primary a');

function closeMenu() {
  if (!header) return;
  header.classList.remove('nav-open');
  burger.setAttribute('aria-expanded', 'false');
  burger.textContent = '☰';
  document.body.style.overflow = '';
}
function openMenu() {
  header.classList.add('nav-open');
  burger.setAttribute('aria-expanded', 'true');
  burger.textContent = '✕';
  document.body.style.overflow = 'hidden';
}
if (burger && header) {
  burger.setAttribute('aria-expanded', 'false');
  burger.addEventListener('click', function () {
    if (header.classList.contains('nav-open')) closeMenu(); else openMenu();
  });
  navLinks.forEach(function (a) {
    a.addEventListener('click', closeMenu);
  });
  document.addEventListener('keydown', function (e) {
    if (e.key === 'Escape') closeMenu();
  });
  window.addEventListener('resize', function () {
    if (window.innerWidth > 860) closeMenu();
  });
}
