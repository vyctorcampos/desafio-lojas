<template>
  <div class="p-6 max-w-4xl mx-auto">
    <h1 class="text-2xl font-bold mb-6">Estabelecimentos</h1>

    <!-- Botões lado a lado -->
    <div class="flex gap-4 mb-6">
      <NuxtLink
        to="/estabelecimentos/novo"
        class="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700"
      >
        + Novo Estabelecimento
      </NuxtLink>

      <NuxtLink
        to="/lojas/novo"
        class="bg-indigo-600 text-white px-4 py-2 rounded hover:bg-indigo-700"
      >
        + Nova Loja
      </NuxtLink>
    </div>

    <div v-if="loading" class="text-gray-500 mb-4">
      Carregando estabelecimentos...
    </div>

    <div v-else>
      <div
        v-for="estabelecimento in estabelecimentos"
        :key="estabelecimento.id"
        class="border p-4 mb-4 rounded shadow-sm flex justify-between items-center"
      >
        <div>
          <h2 class="text-lg font-semibold">
            {{ estabelecimento.numero_estabelecimento }} - {{ estabelecimento.nome }}
          </h2>
          <p class="text-sm text-gray-600 leading-relaxed">
            {{ estabelecimento.razao_social }}<br />
            {{ estabelecimento.endereco }}, {{ estabelecimento.numero }}<br />
            {{ estabelecimento.cidade }} - {{ estabelecimento.estado }} | CEP: {{ estabelecimento.cep }}
          </p>
        </div>

        <div class="flex gap-2">
          <NuxtLink
            :to="`/estabelecimentos/${estabelecimento.id}/lojas`"
            class="bg-blue-500 hover:bg-blue-600 text-white px-3 py-1 rounded"
          >
            Ver Lojas
          </NuxtLink>

          <NuxtLink
            :to="`/estabelecimentos/${estabelecimento.id}`"
            class="bg-yellow-500 hover:bg-yellow-600 text-white px-3 py-1 rounded"
          >
            Editar
          </NuxtLink>

          <button
            @click="removerEstabelecimento(estabelecimento.id)"
            class="bg-red-600 hover:bg-red-700 text-white px-3 py-1 rounded"
          >
            Excluir
          </button>
        </div>
      </div>

      <div v-if="estabelecimentos.length === 0" class="text-gray-600 mt-6">
        Nenhum estabelecimento encontrado.
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useNuxtApp, useRoute } from '#app'

interface Estabelecimento {
  id: number
  numero_estabelecimento: string
  nome: string
  razao_social: string
  endereco: string
  numero: string
  cidade: string
  estado: string
  cep: string
}

interface EstabelecimentoResponse {
  data: Estabelecimento[]
  message: string
  success: boolean
}

const { $api } = useNuxtApp()
const route = useRoute()

const estabelecimentos = ref<Estabelecimento[]>([])
const loading = ref(true)

const carregarEstabelecimentos = async () => {
  loading.value = true
  try {
    const response = await $api<EstabelecimentoResponse>('/estabelecimentos')
    estabelecimentos.value = response.data
  } catch (error) {
    alert('Erro ao carregar estabelecimentos')
    console.error(error)
  } finally {
    loading.value = false
  }
}

onMounted(async () => {
  const novoEstabelecimento = window.history.state?.novoEstabelecimento
  await carregarEstabelecimentos()

  if (
    novoEstabelecimento &&
    !estabelecimentos.value.find((e) => e.id === novoEstabelecimento.id)
  ) {
    estabelecimentos.value.push(novoEstabelecimento)
    window.history.replaceState({}, '', location.href)
  }
})

watch(() => route.fullPath, () => {
  carregarEstabelecimentos()
})

const removerEstabelecimento = async (id: number) => {
  if (!confirm('Deseja realmente excluir este estabelecimento?')) return
  try {
    await $api(`/estabelecimentos/${id}`, { method: 'DELETE' })
    estabelecimentos.value = estabelecimentos.value.filter((e) => e.id !== id)
    alert('Estabelecimento excluído com sucesso!')
  } catch (error) {
    alert('Erro ao excluir estabelecimento')
    console.error(error)
  }
}
</script>
