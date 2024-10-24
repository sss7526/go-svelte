<script context="module">
    export function load({ error, status }) {
      return { props: { error, status } };
    }
  </script>
  
  <script>
    export let error;
    export let status;
  </script>
  
  <main class="flex flex-col items-center justify-center h-screen bg-gray-100">
    <h1 class="text-4xl mb-4">{status}</h1>
    <p class="mb-4">{error.message || 'Page not found.'}</p>
    <a href="/" class="btn btn-primary">Return to Home</a>
  </main>
  
  <style>
    .btn {
      padding: 0.75rem 1.5rem;
      border: none;
      border-radius: 4px;
      cursor: pointer;
      background-color: #007BFF;
      color: white;
      text-decoration: none;
    }
  
    .btn:hover {
      background-color: #0056b3;
    }
  </style>