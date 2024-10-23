<script>
    import '../../app.css';  // Pull in global styles (TailwindCSS, etc.)
  </script>
  
  <!-- Simple layout for login without the Drawer -->
  <main>
    <slot /> <!-- This will render the actual login page -->
  </main>
  
  <style>
    /* Optional: Can apply custom styles for login/signup layout here */
    main {
      display: flex;
      justify-content: center;
      align-items: center;
      height: 100vh;
    }
  </style>