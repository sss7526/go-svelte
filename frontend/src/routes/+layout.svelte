<script>
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
    let open = true; // Control Drawer state
  
    let isAuthenticated = false;
  
    // Function to check authentication status from backend
    async function checkAuth() {
      try {
        const res = await fetch('/api/check-auth', {
          method: 'GET',
          credentials: 'include'
        });
  
        if (res.ok) {
          const session = await res.json();
          isAuthenticated = session.authenticated;
        } else {
          isAuthenticated = false;
        }
  
        if (!isAuthenticated && window.location.pathname !== '/login') {
          goto('/login');
        }
      } catch (err) {
        console.error('Auth check failed', err);
        goto('/login');
      }
    }
  
    onMount(() => {
      checkAuth();
    });
  </script>
  
  {#if isAuthenticated}
    <!-- Drawer layout for authenticated routes -->
    <div class="drawer drawer-mobile">
      <input id="my-drawer" type="checkbox" class="drawer-toggle" bind:checked={open} />
  
      <!-- Main Content Area -->
      <div class="drawer-content">
        <label for="my-drawer" class="btn btn-primary drawer-button">Toggle Sidebar</label>
        <slot />
      </div>
  
      <!-- Sidebar Content -->
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
  
  <style>
    .drawer-content {
      padding: 1rem;
    }
  </style>