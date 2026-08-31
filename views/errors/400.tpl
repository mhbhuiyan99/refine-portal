{{template "layouts/header.tpl" .}}

<section class="error-page">
    <div class="error-content">
        <div class="error-code">400</div>
        <h1>Invalid Request</h1>
        <p>That link looks incomplete or malformed. Please check the URL and try again.</p>
        <a href="/" class="error-home-link">Back to Home</a>
    </div>
</section>

{{template "layouts/footer.tpl" .}}