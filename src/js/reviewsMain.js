document.addEventListener("DOMContentLoaded", function () {
  var el = {
    resultCount: document.getElementById("resultCount"),
    input: document.getElementById("searchInput"),
    clear: document.getElementById("clearBtn"),
    noResults: document.getElementById("noResults"),
    pagination: document.getElementById("pagination"),
    firstBtn: document.getElementById("firstBtn"),
    prevBtn: document.getElementById("prevBtn"),
    nextBtn: document.getElementById("nextBtn"),
    lastBtn: document.getElementById("lastBtn"),
    pageInfo: document.getElementById("pageInfo"),
    container: document.getElementById("reviewContainer"),
  };

  var REACTION_ICONS = {
    heart: "\u2764\uFE0F", thumbs_up: "\uD83D\uDC4D", thumbs_down: "\uD83D\uDC4E",
    laugh: "\uD83D\uDE04", hooray: "\uD83C\uDF89", confused: "\uD83D\uDE15",
    rocket: "\uD83D\uDE80", eyes: "\uD83D\uDC40"
  };

  var GRADIENTS = [
    "linear-gradient(135deg, #667eea 0%, #764ba2 100%)",
    "linear-gradient(135deg, #f093fb 0%, #f5576c 100%)",
    "linear-gradient(135deg, #4facfe 0%, #00f2fe 100%)",
    "linear-gradient(135deg, #43e97b 0%, #38f9d7 100%)",
    "linear-gradient(135deg, #fa709a 0%, #fee140 100%)",
    "linear-gradient(135deg, #a18cd1 0%, #fbc2eb 100%)",
    "linear-gradient(135deg, #fccb90 0%, #d57eeb 100%)",
    "linear-gradient(135deg, #e0c3fc 0%, #8ec5fc 100%)",
    "linear-gradient(135deg, #f5576c 0%, #ff6a88 100%)",
    "linear-gradient(135deg, #667eea 0%, #764ba2 100%)",
    "linear-gradient(135deg, #89f7fe 0%, #66a6ff 100%)",
    "linear-gradient(135deg, #fddb92 0%, #d1fdff 100%)",
    "linear-gradient(135deg, #9890e3 0%, #b1f4cf 100%)",
    "linear-gradient(135deg, #ebc0fd 0%, #d9ded8 100%)",
    "linear-gradient(135deg, #f6d365 0%, #fda085 100%)",
    "linear-gradient(135deg, #fbc2eb 0%, #a6c1ee 100%)",
    "linear-gradient(135deg, #84fab0 0%, #8fd3f4 100%)",
    "linear-gradient(135deg, #cfd9df 0%, #e2ebf0 100%)",
    "linear-gradient(135deg, #a1c4fd 0%, #c2e9fb 100%)",
    "linear-gradient(135deg, #d4fc79 0%, #96e6a1 100%)"
  ];

  function hashString(str) {
    var hash = 0;
    for (var i = 0; i < str.length; i++) {
      hash = ((hash << 5) - hash) + str.charCodeAt(i);
      hash |= 0;
    }
    return Math.abs(hash);
  }

  function applyGradients() {
    var avatars = document.querySelectorAll(".review-avatar[data-author]");
    for (var i = 0; i < avatars.length; i++) {
      var author = avatars[i].dataset.author || "";
      var gradient = GRADIENTS[hashString(author) % GRADIENTS.length];
      avatars[i].style.background = gradient;
    }
  }

  var PAGE_SIZE = 9;
  var allReviews = [];
  var filteredReviews = [];
  var totalPages = 0;
  var currentPage = 1;

  function buildReviewCard(r) {
    var card = document.createElement("div");
    card.className = "review-card";
    card.dataset.url = r.url;

    var reactionsHtml = "";
    if (r.reactions) {
      var keys = Object.keys(r.reactions);
      for (var i = 0; i < keys.length; i++) {
        var k = keys[i];
        if (r.reactions[k] > 0 && REACTION_ICONS[k]) {
          reactionsHtml += '<span class="reaction-chip">' + REACTION_ICONS[k] + " " + r.reactions[k] + "</span>";
        }
      }
    }

    var initial = r.author ? r.author.charAt(0) : "?";
    var gradient = GRADIENTS[hashString(r.author || "") % GRADIENTS.length];

    card.innerHTML =
      '<div class="review-header">' +
        '<div class="review-avatar" style="background:' + gradient + '">' + initial + "</div>" +
        '<div class="review-meta">' +
          '<div class="review-author">' + r.author + "</div>" +
          '<div class="review-timestamp">' + r.timestamp + "</div>" +
        "</div>" +
      "</div>" +
      '<div class="review-body">' + r.body + "</div>" +
      (reactionsHtml ? '<div class="review-reactions">' + reactionsHtml + "</div>" : "");

    return card;
  }

  function clearGrid() {
    el.container.querySelectorAll(".review-card").forEach(function (c) { c.remove(); });
  }

  function renderPage(reviews, page) {
    clearGrid();
    if (page === 1) {
      var header = document.createElement("a");
      header.href = "https://github.com/nisrulz/app-privacy-policy-generator/issues/65";
      header.target = "_blank";
      header.className = "review-card header-card";
      header.innerHTML =
        '<div class="review-body">' +
          '<img src="./downloaded_images/banner.jpg" alt="Banner Image" loading="lazy" />' +
          '<p>Welcome to our virtual Guest Book! Feel free to leave some kind words about App Privacy Policy Generator, share your success stories, or even brag about the big names using our app.</p>' +
          '<p class="click-here">Click on this card to comment</p>' +
        "</div>";
      el.container.appendChild(header);
    }
    var start = (page - 1) * PAGE_SIZE;
    var end = Math.min(start + PAGE_SIZE, reviews.length);
    for (var i = start; i < end; i++) {
      el.container.appendChild(buildReviewCard(reviews[i]));
    }
  }

  function updatePagination() {
    el.pageInfo.textContent = currentPage + " / " + totalPages;
    el.firstBtn.disabled = currentPage <= 1;
    el.prevBtn.disabled = currentPage <= 1;
    el.nextBtn.disabled = currentPage >= totalPages;
    el.lastBtn.disabled = currentPage >= totalPages;
  }

  function showPagination() {
    el.pagination.style.display = "flex";
    updatePagination();
  }

  function hidePagination() { el.pagination.style.display = "none"; }

  function goToPage(page) {
    currentPage = page;
    renderPage(filteredReviews, currentPage);
    el.resultCount.textContent = filteredReviews.length + " reviews";
    el.noResults.style.display = filteredReviews.length === 0 ? "block" : "none";
    showPagination();
  }

  function search(query) {
    var q = query.trim().toLowerCase();
    if (q) {
      filteredReviews = allReviews.filter(function (r) {
        return (r.author + " " + r.body).toLowerCase().includes(q);
      });
    } else {
      filteredReviews = allReviews;
    }
    totalPages = Math.ceil(filteredReviews.length / PAGE_SIZE);
    goToPage(1);
  }

  function clearSearch() {
    el.input.value = "";
    el.clear.style.display = "none";
    filteredReviews = allReviews;
    totalPages = Math.ceil(filteredReviews.length / PAGE_SIZE);
    goToPage(1);
    el.input.focus();
  }

  function debounce(fn, delay) {
    var timer, leading = false;
    return function () {
      var ctx = this, args = arguments;
      if (!leading) { leading = true; fn.apply(ctx, args); }
      clearTimeout(timer);
      timer = setTimeout(function () { leading = false; fn.apply(ctx, args); }, delay);
    };
  }

  function onCardClick(e) {
    var card = e.target.closest(".review-card");
    if (card && card.dataset.url) window.open(card.dataset.url, "_blank");
  }

  function init() {
    document.addEventListener("click", onCardClick);
    el.input.addEventListener("input", debounce(function () {
      el.clear.style.display = el.input.value ? "flex" : "none";
      search(el.input.value);
    }, 300));
    el.clear.addEventListener("click", clearSearch);
    el.input.addEventListener("keydown", function (e) { if (e.key === "Escape") clearSearch(); });
    el.firstBtn.addEventListener("click", function () { goToPage(1); });
    el.prevBtn.addEventListener("click", function () { if (currentPage > 1) goToPage(currentPage - 1); });
    el.nextBtn.addEventListener("click", function () { if (currentPage < totalPages) goToPage(currentPage + 1); });
    el.lastBtn.addEventListener("click", function () { goToPage(totalPages); });

    applyGradients();

    fetch("./reviews-data.json")
      .then(function (res) { return res.json(); })
      .then(function (reviews) {
        el.container.querySelectorAll(".skeleton-card").forEach(function (s) { s.remove(); });
        allReviews = reviews;
        filteredReviews = reviews;
        totalPages = Math.ceil(allReviews.length / PAGE_SIZE);
        el.resultCount.textContent = allReviews.length + " reviews";
        goToPage(1);
        applyGradients();
      })
      .catch(function () {
        el.container.querySelectorAll(".skeleton-card").forEach(function (s) { s.remove(); });
        el.container.insertAdjacentHTML("beforeend", '<div class="loading">Failed to load reviews.</div>');
      });
  }

  init();

  var __themeToggle = window.__themeToggle !== false;
  if (__themeToggle) {
    document.getElementById("themeToggle").style.display = "";
    function updateThemeToggle() {
      var btn = document.getElementById("themeToggle");
      btn.textContent = document.documentElement.getAttribute("data-theme") === "dark" ? "\u2600\uFE0F" : "\uD83C\uDF19";
    }
    function updateThemeLogo() {
      var theme = document.documentElement.getAttribute('data-theme');
      document.querySelectorAll('img[data-theme-logo]').forEach(function (img) {
        var light = img.getAttribute('data-light-src') || img.src;
        if (!img.getAttribute('data-light-src')) {
          img.setAttribute('data-light-src', img.src);
          img.setAttribute('data-dark-src', img.src.replace(/(\.\w+)$/, '_dark$1'));
        }
        img.src = theme === 'dark' ? img.getAttribute('data-dark-src') : img.getAttribute('data-light-src');
      });
    }
    updateThemeToggle();
    updateThemeLogo();
    document.getElementById("themeToggle").addEventListener("click", function () {
      var html = document.documentElement;
      var next = html.getAttribute("data-theme") === "dark" ? "light" : "dark";
      html.setAttribute("data-theme", next);
      localStorage.setItem("theme", next);
      updateThemeToggle();
      updateThemeLogo();
    });
  }
});
