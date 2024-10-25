<script>
  import '../../app.css';  // Import TailwindCSS and DaisyUI
</script>

<main class="flex items-center justify-center h-screen">
  <!-- The login/signup form will render dynamically inside this slot -->
  <slot />
</main>

<style>
  /* Center Content for login/signup pages */
  main {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    /* background-color: #181616; */
  }
</style>