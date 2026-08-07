(function () {
  var API_URL = "https://wcsf-x-academy.onrender.com/api/register";
  var DISCIPLINE_NAMES = { bmx: "BMX Contest Bishkek", workout: "Street Workout Battle", trampoline: "Батут-турнир" };

  var overlay = document.getElementById("register-overlay");
  var form = document.getElementById("register-form");
  var closeBtn = document.getElementById("register-close");
  var statusEl = document.getElementById("register-status");
  var disciplineLabel = document.getElementById("register-discipline-label");
  var currentDiscipline = "";

  function openModal(discipline) {
    currentDiscipline = discipline;
    disciplineLabel.textContent = DISCIPLINE_NAMES[discipline] || "";
    statusEl.textContent = "";
    statusEl.className = "register-status";
    form.reset();
    overlay.classList.add("is-open");
    document.body.style.overflow = "hidden";
  }

  function closeModal() {
    overlay.classList.remove("is-open");
    document.body.style.overflow = "";
  }

  document.querySelectorAll(".js-register-open").forEach(function (btn) {
    btn.addEventListener("click", function () {
      openModal(btn.dataset.discipline);
    });
  });

  closeBtn.addEventListener("click", closeModal);
  overlay.addEventListener("click", function (e) {
    if (e.target === overlay) closeModal();
  });
  document.addEventListener("keydown", function (e) {
    if (e.key === "Escape" && overlay.classList.contains("is-open")) closeModal();
  });

  form.addEventListener("submit", function (e) {
    e.preventDefault();
    var submitBtn = form.querySelector("button[type=submit]");
    var payload = {
      full_name: form.full_name.value.trim(),
      phone: form.phone.value.trim(),
      discipline: currentDiscipline,
      website: form.website.value,
    };

    submitBtn.disabled = true;
    statusEl.className = "register-status";
    statusEl.textContent = "Отправляем...";

    fetch(API_URL, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload),
    })
      .then(function (res) {
        if (!res.ok) throw new Error("bad status");
        statusEl.className = "register-status is-success";
        statusEl.textContent = "Заявка отправлена! Мы свяжемся с вами перед стартом.";
        submitBtn.disabled = false;
        setTimeout(closeModal, 1800);
      })
      .catch(function () {
        statusEl.className = "register-status is-error";
        statusEl.textContent = "Не получилось отправить. Попробуйте ещё раз или напишите нам в WhatsApp.";
        submitBtn.disabled = false;
      });
  });
})();
