import { useContext, useEffect, useState } from "react";
import { AuthContext } from "../firebase_utils/authContext";
import { Link } from "react-router-dom";
import { firebaseAuth } from "../firebase_utils/auth";

export const Home : React.FC = () => {
  const { user } = useContext(AuthContext);
  const [isPending, setIsPending] = useState(false);

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
        {user.email}
      </div>
      <div onClick={signOut}>サインアウト</div>
    </div>
  );

}