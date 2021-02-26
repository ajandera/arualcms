<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-light" id="header">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">arualCMS</a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="float-right" id="navbarSupportedContent">
          <Sign/>
      </div>
    </div>
  </nav>
</template>

<script>

import Sign from '@/components/Sign'

export default {
  name: 'Header',
  components: {
    Sign
  }
}
</script>
