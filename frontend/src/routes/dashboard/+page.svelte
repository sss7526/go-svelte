<script>
    let theme = "white"; // "white" | "g10" | "g80" | "g90" | "g100"
  
    $: document.documentElement.setAttribute("theme", theme);

    import { Tile } from "carbon-components-svelte";

    let message = "Loading...";

    async function loadDashboardMessage() {
        try {
            const res = await fetch("/api/dashboard", {
                method: "GET",
                credentials: "include",
            });

            if (res.status === 401) {
                window.location.href = "/login";
            } else {
                const data = await res.json();
                message = data.message;
            }
        } catch (error) {
            console.error("Failed to load dashboard: ", error);
        }
    } 

    loadDashboardMessage();
</script>

<div class="dashboard-container">
    <Tile>
        <h2>Dashboard</h2>
        <p>{message}</p>
    </Tile>
</div>

<style>
    .dashboard-container {
        padding: 2rem;
        margin-top: 50px;
        max-width: 500px;
        margin: 50px auto;
    }
</style>