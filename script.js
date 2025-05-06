window.addEventListener("load", () => {
    fetchNews("India"); // default news on load
});

async function fetchNews(query) {
    try {
        const res = await fetch(`/search?q=${encodeURIComponent(query)}`);
        const data = await res.json();
        bindData(data.articles);
    } catch (error) {
        console.error('Error fetching news:', error);
    }
}

function bindData(articles) {
    const container = document.getElementById("news-container");
    const template = document.getElementById("news-template");
    container.innerHTML = "";

    articles.forEach(article => {
        if (!article.urlToImage) return;

        const clone = template.content.cloneNode(true);
        const card = clone.querySelector(".news-card");

        clone.querySelector("#news-img").src = article.urlToImage;
        clone.querySelector("#news-title").innerText = article.title;
        clone.querySelector("#news-desc").innerText = article.description;
        clone.querySelector("#news-source").innerText =
            `${article.source.name} · ${new Date(article.publishedAt).toLocaleString()}`;

        // Enable card click
        card.addEventListener("click", () => {
            window.open(article.url, "_blank");
        });

        container.appendChild(clone);
    });
}

const searchButton = document.getElementById("search-button");
const searchText = document.getElementById("search-text");

searchButton.addEventListener("click", () => {
    const query = searchText.value.trim();
    if (!query) return;
    fetchNews(query);
});
