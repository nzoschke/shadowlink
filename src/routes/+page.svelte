<script lang="ts">
  import { onMount } from "svelte";
  import { createClient } from "@supabase/supabase-js";
  import type { Database } from "$lib/database";
  import { relative } from "$lib/time";

  type Item = {
    author: Author;
    meta: Meta;
    url: string;
    updated_at: string;
  };

  type Author = {
    channel: string;
    name: string;
    service: string;
  };

  type Meta = {
    url: string;
    html: string;
    type: string;
    title: string;
    width: number;
    height: number;
    author_url: string;
    author_name: string;
    description: string;
    provider_url: string;
    provider_name: string;
    thumbnail_url: string;
    thumbnail_width: number;
    thumbnail_height: number;
  };

  const supabase = createClient<Database>(
    "https://fdbwwvmqzpgoegjbhefz.supabase.co",
    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImZkYnd3dm1xenBnb2VnamJoZWZ6Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3MzM0NTQ5MDMsImV4cCI6MjA0OTAzMDkwM30.foyyRENh5S1VSWKMIRmExvAIePcfAJUlSeCDVB9_9WQ",
  );

  let items = $state<Item[]>([]);

  onMount(async () => {
    const { data, error } = await supabase
      .from("items")
      .select("*")
      .is("deleted_at", null)
      .limit(100);
    if (error) {
      console.error(error);
      return;
    }

    items = data as any as Item[];
  });
</script>

<svelte:head>
  <title>Shadow Link</title>
</svelte:head>

<div class="flex size-full flex-col items-center justify-center">
  <div class="navbar bg-base-300">
    <div class="flex-1">
      <div class="text-xl">Discord Community Playlist</div>
    </div>
    <div class="flex-none gap-2">
      <a href="https://github.com/nzoschke/shadowlink" class="btn btn-square btn-ghost btn-sm">
        <img alt="github" src="/github.svg" />
      </a>
    </div>
  </div>

  <div class="w-full space-y-2 p-2">
    {#each items as _, i}
      {@render item(items[items.length - 1 - i])}
    {/each}
  </div>
</div>

{#snippet item(i: Item)}
  {@const { author, meta, updated_at, url } = i}

  <div class="flex size-full space-x-2 rounded border bg-base-200">
    <a class="avatar size-16" href={i.url}>
      <div class="rounded">
        {#if meta.thumbnail_url}
          <img src={meta.thumbnail_url} alt="" />
        {/if}
      </div>
    </a>
    <div class="flex flex-1 flex-col justify-between overflow-hidden">
      <div class="truncate">
        <a class="text-lg font-bold" class:italic={!meta.title} href={url}
          >{meta.title || "no title"}</a
        >
      </div>
      <div class="truncate">
        by <b>{author.name}</b>
      </div>
    </div>
    <div class="flex-0 flex w-40 flex-col justify-between overflow-hidden text-sm">
      <div class="truncate">
        <a class="leading-7" href={url}>{url.split("/")[2]}</a>
      </div>
      <div class="truncate">{relative(Date.parse(updated_at))}</div>
    </div>
  </div>
{/snippet}
