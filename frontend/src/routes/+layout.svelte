<script>
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
    import { page } from '$app/stores';
    // import { page } from '$app/stores'; // Store to watch changes in route state
  
    // const authRouteExceptions = ['/login', '/signup'];  // Routes that don't need auth
  
    // Mock authenticated state (fetched from actual Go backend API)
    let isAuthenticated = false;
    let authChecked = true;
    let open = true; // Drawer open/close state

    $: currentPath = $page.url.pathname;

    let errorMessage;
  
    // Function to check authentication status via API
    async function checkAuth() {
      try {
        const res = await fetch('/api/check-auth', {
          method: 'GET',
          credentials: 'include',  // Always include session/cookies
        });
  
        if (res.ok) {
          const session = await res.json();
          isAuthenticated = session.authenticated;
        } else {
          isAuthenticated = false;
          throw new Error("Unauthorized");
        }
  
        // // Redirect to /login if not authenticated and if not on a public page
        // if (!isAuthenticated && !authRouteExceptions.includes(page.url.pathname)) {
        //   goto('/login');
        // }
      } catch (error) {

        console.error("check auth failed:", error);
        errorMessage = error.errorMessage;
        isAuthenticated = false;
      } finally {

        authChecked = true;
        // console.error('Authentication check failed:', error);
        if (!isAuthenticated && window.location.pathname !== '/login') {
          goto('/login');
        }
      }
    }
  
    onMount(() => {
      checkAuth();  // Run auth check once page loads.
    });

    import '../app.css';
  </script>

  <slot />

  {#if errorMessage}
    <p class="text-red-500">{errorMessage}</p>
  {/if}
  
  {#if authChecked}
    {#if isAuthenticated}
      <!-- Drawer layout starts when user is authenticated -->
      <div class="drawer drawer-mobile">
        <input id="my-drawer" type="checkbox" class="drawer-toggle" bind:checked={open}>
    
        <div class="drawer-content">
          <label for="my-drawer" class="btn btn-primary drawer-button">Toggle Sidebar</label>
          <slot />  <!-- The main content for the current route -->
        </div>
    
        <div class="drawer-side">
          <label for="my-drawer" class="drawer-overlay"></label> 
          <ul class="menu p-4 w-80 bg-base-100 text-base-content">
            <li><a href="/dashboard/channel1">Channel 1</a></li>
            <li><a href="/dashboard/channel2">Channel 2</a></li>
            <li><a href="/settings">Settings</a></li>
          </ul>
        </div>
      </div>
    {/if}

    {#if !isAuthenticated && currentPath === '/login'}
      <slot />
    {/if}

  {/if}

  
  
  <style>
    /* Simple styling for the content inside Drawer */
    .drawer-content {
      padding: 1rem;
    }
  </style>