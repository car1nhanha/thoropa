<script setup lang="ts">
import { faDownload } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { computed, ref, watchEffect } from "vue";
import DownloadCsv from "vue-json-csv";
import { ApiService, type ResponseGetLink } from "../../../service/api";
import Button from "../../atoms/button.vue";
import Input from "../../molecules/Input.vue";
import History from "../../organisms/history.vue";

const appURL = import.meta.env.VITE_APP_URL;
const appName = import.meta.env.VITE_APP_NAME;

const links = ref<ResponseGetLink[]>([]);
const input = ref("");
const shortenedLink = ref("");

watchEffect(() => {
  ApiService.getLinks().then((response) => {
    links.value = response;
  });
});

const inputIsValid = computed(() => {
  return /^https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&\/=]*)$/.test(
    input.value,
  );
});

const csvData = computed(() => {
  return links.value.map((link) => ({
    id: link.Id,
    original: link.Original,
    short: `${appURL}/${link.Id}`,
    accesses: link.Accesses,
    created_at: new Date(link.CreatedAt * 1000).toISOString(),
  }));
});

const shortenLink = () => {
  if (!inputIsValid.value) return;

  ApiService.createLink(input.value).then((response) => {
    shortenedLink.value = `${appURL}/${response.Id}`;
    links.value.unshift(response);
  });
};

const reloadLinks = () => {
  ApiService.getLinks().then((response) => {
    links.value = response;
  });
};
</script>

<template>
  <div class="content">
    <div class="header">
      <h1>{{ appName }}</h1>
    </div>
    <div class="main">
      <div class="form">
        <h2 class="title">New link</h2>

        <div class="form-content">
          <Input label="Original link" placeholder="www.example.com" v-model="input" />
          <Input label="shortened link" :placeholder="`${appURL}/`" :disabled="true" v-model="shortenedLink" />
        </div>

        <Button size="large" @click="shortenLink" :disabled="!inputIsValid">Shorten</Button>
      </div>

      <div class="history">
        <div class="history-header">
          <h2 class="title">My links</h2>
          <DownloadCsv :data="csvData" name="links.csv">
            <Button size="small" :disabled="!links.length">
              <template #icon>
                <FontAwesomeIcon :icon="faDownload" />
              </template>
              Download CSV
            </Button>
          </DownloadCsv>
        </div>

        <History :links="links" :reloadLinks="reloadLinks" />
      </div>
    </div>
  </div>
</template>

<style src="./home.css" lang="css" scoped></style>
