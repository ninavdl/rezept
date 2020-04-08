const fieldMarker = "data-rezept-dynamic";

function setTag(name: string, content: string) {
    const el = document.createElement("meta");
    el.name = name;
    el.content = content;
    el.setAttribute(fieldMarker, "");
    document.head.appendChild(el);
}

export function setMetadata(pageTitle: string, metadata: Metadata) { 
    let title = pageTitle;
    if (metadata.title != "") title = metadata.title + " | " + title;
    document.title = title;
    Array.from(document.querySelectorAll(`meta[${fieldMarker}]`)).forEach(tag => tag.parentNode.removeChild(tag));

    setTag("og:title", title);
    setTag("og:site_name", pageTitle);
    setTag("twitter:card", "summery_large_image");
    setTag("twitter:title", title);

    if (metadata.description != null) {
        setTag("description", metadata.description);
        setTag("og:description", metadata.description);
        setTag("twitter:description", metadata.description);
    }
    if (metadata.author != null) {
        setTag("author", metadata.author);
    }
    if (metadata.date != null) {
        setTag("date", metadata.date.toISOString());
    }
    if (metadata.tags != null) {
        setTag("keywords", metadata.tags.join(", "));
    }
    if (metadata.imageURL != null) {
        setTag("og:image", metadata.imageURL);
        setTag("twitter:image", metadata.imageURL);
    }
}

export class Metadata {
    title: string;
    description: string;
    author: string;
    date: Date;
    tags: string[];
    imageURL: string;
}