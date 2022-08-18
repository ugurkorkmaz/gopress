<script>
  import { page } from "$app/stores";
</script>

<section class="error">
  <div class="main">
    <h1 class="header">{$page.status}</h1>
    <p class="status">{$page.error.message}</p>
    <p class="message">
      Sorry, we can't find that page. You'll find lots to explore on the home
      page.
    </p>
    <a class="back" href="/">Back to Homepage</a>
  </div>
</section>

<style>
  .error {
    @apply bg-white dark:bg-gray-900;
  }
  .main {
    @apply py-8 px-4 mx-auto max-w-screen-xl lg:py-16 lg:px-6 text-center;
  }
  .error .main .header {
    @apply mb-4 text-7xl tracking-tight font-extrabold lg:text-9xl text-blue-600 dark:text-blue-500;
  }
  .error .main .back {
    @apply inline-flex text-white bg-blue-600 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:focus:ring-blue-900 my-4;
  }
  .error .main .status {
    @apply mb-4 text-3xl tracking-tight font-bold text-gray-900 md:text-4xl dark:text-white;
  }
  .error .main .message {
    @apply mb-4 text-lg font-light text-gray-500 dark:text-gray-400;
  }
</style>
