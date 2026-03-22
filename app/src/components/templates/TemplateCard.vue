<script lang="ts" setup>
const props = defineProps({
  title: {
    type: String,
    required: true,
  },
});
</script>

<template>
  <div class="template-card">
    <div class="template-card-body">
      <slot name="image"></slot>
      <h1 class="title">{{ props.title }}</h1>
      <p class="description">
        <slot></slot>
      </p>
    </div>
  </div>
</template>

<style lang="css" scoped>
.template-card {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100vh;
  min-height: 600px;
}

.template-card-body {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 64px 48px;
  gap: 24px;

  width: 580px;
  background: var(--gray-100);
  border-radius: 8px;
}

.template-card-body img {
  width: 48px;
  height: 48px;
}

.template-card-body h1 {
  font-weight: 700;
  font-size: 24px;
  line-height: 32px;
  color: var(--gray-600);
}

.description {
  font-weight: 600;
  font-size: 14px;
  line-height: 18px;
  color: var(--gray-500);
  text-align: center;
}

.description a {
  color: var(--blue-base);
  text-decoration: none;
}
</style>
