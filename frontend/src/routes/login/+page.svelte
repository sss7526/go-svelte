<script>
    import { TextInput } from "carbon-components-svelte";
    import { Button } from "carbon-components-svelte";

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


<div class="bx--form-item">
    {#if errorMessage !== ""}
        <p class="bx--label--danger">{errorMessage}</p>
    {/if}

    <form on:submit={handleLogin}>
        <TextInput
            bind:value={username}
            id="username"
            labelText="username"
            placeholder="Enter username"
        />
        <TextInput
            bind:value={password}
            id="password"
            type="password"
            labelText="Password"
            placeholder="Enter password"
        />
        <Button kind="primary" type="submit">Login</Button>
    </form>
</div>

<style>
    .bx--form-item {
        max-width: 400px;
        margin: 100px auto;
    }
</style>