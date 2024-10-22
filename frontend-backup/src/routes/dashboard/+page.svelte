<script>

    import Card from "@smui/card";

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


<Card variant="outlined" style="margin: 50px auto; max-width: 600px;">
    <h2>Dashboard</h2>
    <p>{message}</p>
</Card>


<style>
    h2 {
        margin-bottom: 1rem;
    }
</style>