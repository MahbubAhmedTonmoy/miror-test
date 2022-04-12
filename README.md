# miror-test

$ git clone --mirror $URL


is a short-hand for

$ git clone --bare $URL
$ (cd $(basename $URL) && git remote add --mirror=fetch origin $URL)

$------------------
$cd __.git
$git push --mirror {new repo URL}
