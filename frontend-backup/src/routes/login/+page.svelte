<script>
    import Button from "@smui/button";
    import TextField from "@smui/textfield";
    import FormField from "@smui/form-field"

    let username = "";
    let password = "";
    let errorMessage = "";

    async function handleLogin(event) {
        event.preventDefault();

        const response = await fetch("/api/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                username,
                password,
            }),
        });

        if (response.ok) {
            window.location.href = "/dashboard";
        } else {
            errorMessage = "Invalid credentials";
        }
    }
</script>



<div>
    {#if errorMessage !== ""}
        <p class="error">{errorMessage}</p>
    {/if}

    <form on:submit={handleLogin}>
        <FormField>
            <TextField
            bind:value={username}
            label="username"
            variant="outlined"
            required
        />
        </FormField>
            <TextField
            type="password"
            bind:value={password}
            label="Password"
            variant="outlined"
            required
        />

        <Button color="primary" type="submit" variant="raised">
            Login
        </Button>
    </form>
</div>

<style>
    .error {
        color: red;
        margin-bottom: 1rem;
    }

    form {
        max-width: 400px;
        margin: 100px auto;
        display: flex;
        flex-direction: column;
        justify-content: center;
        gap: 1rem;
        /* width: 100%; */
    }
</style>