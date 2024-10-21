<script>
    let isAuthenticated = false;

    async function checkAuthStatus() {
        const response = await fetch("/api/check-auth", { method: "GET" });

        if (response.status === 200) {
            isAuthenticated = true;
            window.location.href = "/dashboard";
        } else {
            window.location.href = "/login";
        }
    }

    checkAuthStatus();
</script>

<div>
    <p>Redirection...</p>
</div>