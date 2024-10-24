<script>
  import '../../app.css'; // Correct path if necessary
</script>

<main class="flex items-center justify-center h-screen bg-gray-100">
  <slot />
</main>

<style>
  main {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-color: #f0f0f0;
  }
</style>