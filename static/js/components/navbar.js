function renderNavbar() {

    const navbar = document.getElementById("navbar");

    navbar.innerHTML = `
        <header class="od-navbar">

            <div class="nav-left">

                <a href="#">FIND A RENTAL</a>

                <a href="#">DESTINATIONS</a>

                <a href="#">ABOUT</a>

            </div>

            <div class="nav-center">

                <a href="/" class="logo">

                    <img
                        src="/static/images/logo.png"
                        alt="OwnerDirect">

                </a>

            </div>

            <div class="nav-right">

                <button
                    class="menu-btn"
                    type="button">

                    <span></span>
                    <span></span>
                    <span></span>

                </button>

            </div>

        </header>
    `;
}