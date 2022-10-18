import urlJoin from "url-join";

const baseUrl = import.meta.env.VITE_API_BASE_URL

const apiUrls = {
  storageItemList: urlJoin(baseUrl, '/index'),
}

type StorageItemList = {
  itemNames: string[],
}

function getStorageItemList (uid : string) {
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