module Deque (Deque, mkDeque, pop, push, shift, unshift) where

import Data.IORef

data Deque a = Deque { dqSize :: IORef Integer
                     , dqFirst :: IORef (Maybe (Node a))
                     , dqLast :: IORef (Maybe (Node a))
                     }

data Node a = Node { nValue :: a
                   , nPrev :: IORef (Maybe (Node a))
                   , nPost :: IORef (Maybe (Node a))
                   }

mkDeque :: IO (Deque a)
mkDeque = do
  size <- newIORef 0
  first <- newIORef Nothing
  last <- newIORef Nothing
  pure (Deque { dqSize = size
              , dqFirst = first
              , dqLast = last
              })

pop :: Deque a -> IO (Maybe a)
pop deque@Deque { dqSize = dqSizeRef
                , dqFirst = dqFirstRef
                , dqLast = dqLastRef
                } = do
  maybeLast <- readIORef dqLastRef
  case maybeLast of
    Nothing -> pure Nothing
    Just poppedNode@Node { nValue = poppedValue
                         , nPrev = newLastRef
                         } -> do
      modifyIORef' dqSizeRef (+ (-1))
      newLastMaybeNode <- readIORef newLastRef
      case newLastMaybeNode of
        Nothing -> do
          writeIORef dqFirstRef Nothing
          writeIORef dqLastRef Nothing
        Just newLastNode@Node { nPost = newLastNodePostRef } -> do
          writeIORef newLastNodePostRef Nothing
          writeIORef dqLastRef (Just newLastNode)
      pure (Just poppedValue)

push :: Deque a -> a -> IO ()
push deque@Deque { dqSize = dqSizeRef
                 , dqFirst = dqFirstRef
                 , dqLast = dqLastRef
                 } x = do
  modifyIORef' dqSizeRef (+ 1)
  maybeLast <- readIORef dqLastRef
  prevNode <- newIORef maybeLast
  postNode <- newIORef Nothing
  let newNode = Node { nValue = x
                     , nPrev = prevNode
                     , nPost = postNode
                     }
  case maybeLast of
    Nothing -> do
      writeIORef dqFirstRef (Just newNode)
      writeIORef dqLastRef (Just newNode)
    Just oldLastNode@Node { nPost = oldLastNodePostRef } -> do
      writeIORef dqLastRef (Just newNode)
      writeIORef oldLastNodePostRef (Just newNode)

unshift :: Deque a -> a -> IO ()
unshift deque@Deque { dqSize = dqSizeRef
                    , dqFirst = dqFirstRef
                    , dqLast = dqLastRef
                    } x = do
  modifyIORef' dqSizeRef (+ 1)
  maybeFirst <- readIORef dqFirstRef
  prevNode <- newIORef Nothing
  postNode <- newIORef maybeFirst
  let newNode = Node { nValue = x
                     , nPrev = prevNode
                     , nPost = postNode
                     }
  case maybeFirst of
    Nothing -> do
      writeIORef dqFirstRef (Just newNode)
      writeIORef dqLastRef (Just newNode)
    Just oldFirstNode@Node { nPrev = oldFirstNodePrevRef } -> do
      writeIORef dqFirstRef (Just newNode)
      writeIORef oldFirstNodePrevRef (Just newNode)

shift :: Deque a -> IO (Maybe a)
shift deque@Deque { dqSize = dqSizeRef
                , dqFirst = dqFirstRef
                , dqLast = dqLastRef
                } = do
  maybeFirst <- readIORef dqFirstRef
  case maybeFirst of
    Nothing -> pure Nothing
    Just poppedNode@Node { nValue = poppedValue
                         , nPost = newFirstRef
                         } -> do
      modifyIORef' dqSizeRef (+ (-1))
      newFirstMaybeNode <- readIORef newFirstRef
      case newFirstMaybeNode of
        Nothing -> do
          writeIORef dqFirstRef Nothing
          writeIORef dqLastRef Nothing
        Just newFirstNode@Node { nPrev = newFirstNodePrevRef } -> do
          writeIORef newFirstNodePrevRef Nothing
          writeIORef dqFirstRef (Just newFirstNode)
      pure (Just poppedValue)