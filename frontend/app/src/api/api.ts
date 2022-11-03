import urlJoin from "url-join";

const baseUrl = import.meta.env.VITE_API_BASE_URL

const apiUrls = {
  storageItemList: urlJoin(baseUrl, '/index'),
}

export type StorageItemNames = {
  itemNames: string[],
}

function getStorageItemList (uid : string) : Promise<Response> {
  const params = new URLSearchParams();
  params.append("uid", uid);
  return fetch(apiUrls.storageItemList, 
    {
      method: 'post',
      body: params,
    }
  );
}

export const ApiClient = {
  getStorageItemList: getStorageItemList,
}