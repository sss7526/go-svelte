<script>
	// Import Tailwind styles via app.css
	import '../app.css';
	
	// State variable to control whether the drawer (sidebar) is open or closed
	let open = true;
  </script>
  
  <!-- Sidebar Layout with DaisyUI Drawer component -->
  <div class="drawer drawer-mobile">
	<!-- Control checkbox for toggling sidebar (bound to 'open' state) -->
	<input id="my-drawer" type="checkbox" class="drawer-toggle" bind:checked={open}>
	
	<!-- Main content wrapper for routed pages -->
	<div class="drawer-content">
	  <!-- Button to toggle the sidebar -->
	  <label for="my-drawer" class="btn btn-primary drawer-button">Toggle Sidebar</label>
	  <!-- Inject dynamic route content here -->
	  <slot />
	</div>
  
	<!-- Sidebar (Drawer) content -->
	<div class="drawer-side">
	  <label for="my-drawer" class="drawer-overlay"></label>
	  <ul class="menu p-4 w-80 bg-base-100 text-base-content">
		<li><a href="#channel1">Channel 1</a></li>
		<li><a href="#channel2">Channel 2</a></li>
		<li><a href="#settings">Settings</a></li>
	  </ul>
	</div>
  </div>
  
  <style>
	/* Minimal layout padding for content */
	.drawer-content {
	  padding: 1rem;
	}
  </style>
