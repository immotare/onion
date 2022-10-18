import { useContext, useEffect, useState } from "react";
import { AuthContext } from "../firebase_utils/authContext";
import { Link } from "react-router-dom";
import { firebaseAuth } from "../firebase_utils/auth";
import { ApiClient } from "../api/api";

export const Home : React.FC = () => {
  const { user } = useContext(AuthContext);
  const [isPending, setIsPending] = useState(false);
  const [storageItems, setStorageItems] = useState<string[]>([]);

  useEffect(() => {
    if (user) {
      console.log(user.uid);
      ApiClient.getStorageItemList(user.uid)
        .then((response) => {
          console.log(response);
        })
        .catch((reason) => {
          console.log(reason);
        });
    }
  }, [user]);

  if (!user) {
    return (
      <div>
        <div ><Link to="signin">サインイン</Link></div>
        <div><Link to="create_account">アカウントを作成</Link></div>
      </div>
    );
  }

  const signOut = () => {
    if (!isPending) {
      setIsPending(true);
      firebaseAuth.signOut()
        .then(() => {
          setIsPending(false);
        })
        .catch((reason) => {
          console.log(reason);
        });
    }
  }

  return (
    <div>
      <div>
        {storageItems.map((name, idx) => {
          return (
            <div key={idx}>
              {name}
            </div>
          );
        })}
      </div>
      <div onClick={signOut}>サインアウト</div>
    </div>
  );

}