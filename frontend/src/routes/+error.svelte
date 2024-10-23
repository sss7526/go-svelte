<script context="module">
    export function load({ error, status }) {
      return { props: { error, status } };
    }
  </script>
  
  <script>
    export let error;
    export let status;
  </script>
  
  <main class="error-page">
    <h1>{status}</h1> <!-- Display error code like 404 -->
    <p>{error.message || 'Something went wrong.'}</p>
  
    <a href="/" class="btn btn-primary">Return to Home</a> <!-- Back to Home button -->
  </main>
  
  <!-- Styling for the 404 page -->
  <style>
    .error-page {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      height: 100vh;
      text-align: center;
    }
  
    h1 {
      font-size: 4rem;
    }
  
    p {
      margin-top: 1.5rem;
      font-size: 1.25rem;
    }
  
    a {
      margin-top: 2rem;
    }
  </style>