<script lang="ts">
    import { onMount } from 'svelte';  // Lifecycle hook
    let isSignUp = false;  // Track whether the user is in 'login' or 'signup' mode
  
    let username = '';  // Username field
    let password = '';  // Password field
    let errorMessage = '';  // Handle errors
  
    // Function to toggle between login and signup forms
    function toggleMode() {
      isSignUp = !isSignUp;
    }
  
    // Function to handle form submission (login or signup based on mode)
    async function handleFormSubmit(e: Event) {
      e.preventDefault();
  
      const endpoint = isSignUp ? '/api/signup' : '/api/login';  // Switch between login/signup API endpoints
  
      const response = await fetch(endpoint, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
      });
  
      if (response.ok) {
        window.location.href = '/dashboard';  // Redirect to dashboard on success
      } else {
        const error = await response.text();  // Show error if user can't sign in
        errorMessage = error || 'Something went wrong!';
      }
    }
  </script>
  
  <!-- UI for both login and signup -->
  <div class="form-container">
    <h2>{isSignUp ? 'Sign Up' : 'Login'}</h2>
  
    {#if errorMessage !== ''}
      <p class="text-red-500">{errorMessage}</p>  <!-- Show error messages dynamically -->
    {/if}
  
    <form on:submit={handleFormSubmit} class="form">
      <!-- Username and password inputs -->
      <input type="text" bind:value={username} placeholder="Username" required class="input input-bordered" />
      <input type="password" bind:value={password} placeholder="Password" required class="input input-bordered" />
  
      <button type="submit" class="btn btn-primary">{isSignUp ? 'Sign Up' : 'Login'}</button>
    </form>
  
    <!-- Link to switch between login and signup -->
    <p class="mt-3">
      {#if isSignUp}
        Already have an account? <a href="javascript:void(0);" on:click={toggleMode} class="link">Log in</a>
      {:else}
        Don’t have an account? <a href="javascript:void(0);" on:click={toggleMode} class="link">Sign up</a>
      {/if}
    </p>
  </div>
  
  <style>
    /* Minor layout and spacing tweaks for neat UI */
    .form-container {
      max-width: 400px;
      margin: 50px auto;
      text-align: center;
    }
  
    input {
      margin: 10px 0;
      padding: 10px;
      width: 100%;
    }
  
    button {
      margin-top: 15px;
      width: 100%;
    }
  
    .link {
      color: #007BFF;
      cursor: pointer;
      text-decoration: underline;
    }
  </style>