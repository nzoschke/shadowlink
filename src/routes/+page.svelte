<script lang="ts">
  import { onMount } from "svelte";
  import { createClient } from "@supabase/supabase-js";
  import type { Database } from "$lib/database";
  import { relative } from "$lib/time";

  type Item = Database["public"]["Tables"]["items"]["Row"];
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
    const { data, error } = await supabase.from("items").select("*").limit(10);
    if (error) {
      console.error(error);
      return;
    }

    items = data;
  });
</script>

<svelte:head>
  <title>Shadow Link</title>
</svelte:head>

<div class="flex size-full flex-col items-center justify-center gap-2 p-2">
  {#each items as _, i}
    {@const item = items[items.length - 1 - i]}
    {@const meta = item.meta as Meta}

    <div class="flex size-full space-x-2 rounded border bg-base-200">
      <a class="avatar size-16" href={item.link}>
        <div class="rounded">
          {#if meta.thumbnail_url}
            <img src={meta.thumbnail_url} alt="" />
          {/if}
        </div>
      </a>
      <div class="flex flex-1 flex-col justify-between overflow-hidden">
        <div class="truncate">
          <a class="text-lg font-bold" class:italic={!meta.title} href={item.link}
            >{meta.title || "no title"}</a
          >
        </div>
        <div class="truncate">
          by <b>{item.user_name}</b> in {item.channel_name}
        </div>
      </div>
      <div class="flex-0 flex w-40 flex-col justify-between overflow-hidden text-sm">
        <div class="truncate">
          <a class="leading-7" href={item.link}>{item.link.split("/")[2]}</a>
        </div>
        <div class="truncate">{relative(Date.parse(item.created_at))}</div>
      </div>
    </div>
  {/each}
</div>
