{{template "layouts/header.tpl" .}}

<section class="error-page">
    <div class="error-content">
        <div class="error-code">500</div>
        <h1>Something Went Wrong</h1>
        <p>We hit an unexpected error on our end. Please try again in a moment.</p>
        <a href="/" class="error-home-link">Back to Home</a>
    </div>
</section>

{{template "layouts/footer.tpl" .}}