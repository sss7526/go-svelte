import forms from '@tailwindcss/forms';
import aspectRatio from '@tailwindcss/aspect-ratio';
import daisyui from 'daisyui';  // Import DaisyUI plugin

/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './src/**/*.{html,js,svelte,ts}',  // Ensures Tailwind scans your Svelte and JS code
    './src/routes/**/*.{html,js,svelte,ts}' // Scans route files
  ],  

  theme: {
    extend: {},  // Extend theme if needed later for customization
  },

  plugins: [
    forms,          // Plugin for better forms styling
    aspectRatio,    // Plugin for controlling aspect ratios in images or containers
    daisyui         // Add DaisyUI for pre-built TailwindCSS component library
  ],
};