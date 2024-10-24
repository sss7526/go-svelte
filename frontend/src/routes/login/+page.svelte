<script>
  let isSignUp = false; // Track mode

  let username = '';
  let password = '';
  let errorMessage = '';

  // Toggle between Login and Signup
  function toggleMode() {
      isSignUp = !isSignUp;
      errorMessage = ''; // Reset error message on mode switch
  }

  // Handle form submission
  async function handleFormSubmit(e) {
      e.preventDefault();

      const endpoint = isSignUp ? '/api/signup' : '/api/login';

      const response = await fetch(endpoint, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify({ username, password }),
      });

      if (response.ok) {
          window.location.href = '/dashboard'; // Redirect on success
      } else {
          const error = await response.text();
          errorMessage = error || 'Something went wrong!';
      }
  }
</script>

<div class="w-full max-w-md p-8 bg-white rounded shadow">
  <h2 class="text-2xl mb-4">{isSignUp ? 'Sign Up' : 'Login'}</h2>

  {#if errorMessage}
      <p class="text-red-500 mb-4">{errorMessage}</p>
  {/if}

  <form on:submit={handleFormSubmit} class="flex flex-col gap-4">
      <input
          type="text"
          bind:value={username}
          placeholder="Username"
          required
          class="input input-bordered"
      />
      <input
          type="password"
          bind:value={password}
          placeholder="Password"
          required
          class="input input-bordered"
      />
      <button type="submit" class="btn btn-primary">
          {isSignUp ? 'Sign Up' : 'Login'}
      </button>
  </form>

  <p class="mt-4 text-center">
      {#if isSignUp}
          Already have an account? <a href="javascript:void(0);" on:click={toggleMode} class="link">Log in</a>
      {:else}
          Don’t have an account? <a href="javascript:void(0);" on:click={toggleMode} class="link">Sign up</a>
      {/if}
  </p>
</div>

<style>
  .input {
      width: 100%;
      padding: 0.5rem;
      border: 1px solid #ccc;
      border-radius: 4px;
  }

  .btn {
      padding: 0.75rem;
      border: none;
      border-radius: 4px;
      cursor: pointer;
  }

  .btn-primary {
      background-color: #007BFF;
      color: white;
  }

  .link {
      color: #007BFF;
      text-decoration: underline;
      cursor: pointer;
  }

  .link:hover {
      text-decoration: none;
  }
</style>