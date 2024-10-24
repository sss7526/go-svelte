<script>
    import { onMount } from 'svelte';
  
    let message = "Loading...";
  
    async function loadDashboardMessage() {
      const res = await fetch('/api/dashboard', {
        method: 'GET',
        credentials: 'include'
      });
  
      if (res.ok) {
        const data = await res.json();
        message = data.message;
      } else {
        window.location.href = '/login'; // Redirect to login if unauthorized
      }
    }
  
    onMount(() => {
      loadDashboardMessage();
    });
  </script>
  
  <div class="max-w-4xl mx-auto p-4">
    <h1 class="text-3xl mb-4">Dashboard</h1>
    <p>{message}</p>
  </div>
  
  <style>
    /* Optional additional styles */
  </style>