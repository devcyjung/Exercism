       IDENTIFICATION DIVISION.
       PROGRAM-ID. BINARY-SEARCH.

       DATA DIVISION.
       WORKING-STORAGE SECTION.
       01 WS-FOUND-NOT-FOUND     PIC 9 VALUE 0.
           88 FOUND              VALUE 1.
       01 WS-ITEM                PIC 9999.
       01 WS-RESULT              PIC 99 VALUE 0.
       01 WS-ERROR               PIC X(40) VALUE SPACES.
       01 WS-COUNT               PIC 99 VALUE 0.
       01 WS-ARRAY               PIC X(20).  
       01 WS-SUBSTR              PIC X(4).
       01 WS-PTR                 PIC 99 VALUE 1.
       01 WS-TABLE.
           02 WS-ELEMENT
               OCCURS 1 TO 20 DEPENDING ON WS-COUNT
               INDEXED BY IDX
               PIC 9(4).
       01 WS-FIRST-INDEX         PIC 99.
       01 WS-MIDDLE-INDEX        PIC 99.
       01 WS-LAST-INDEX          PIC 99.
       01 WS-MIDDLE-ELEMENT      PIC 9999.
       01 WS-TEMP                PIC 9999.
       01 WS-LOOP-IDX            PIC 99.

       PROCEDURE DIVISION.
       BINARY-SEARCH.
           INITIALIZE WS-FOUND-NOT-FOUND WS-RESULT WS-ERROR
                      WS-PTR WS-COUNT WS-FIRST-INDEX
                      WS-MIDDLE-INDEX WS-LAST-INDEX.

           PERFORM PARSE-ARRAY.

           IF WS-COUNT = 0
               MOVE "value not in array" TO WS-ERROR
               GO TO DISPLAY-RESULT
           END-IF.

           DISPLAY "find " WS-ITEM " in array".

           MOVE 1 TO WS-FIRST-INDEX.
           MOVE WS-COUNT TO WS-LAST-INDEX.

           PERFORM UNTIL WS-FIRST-INDEX > WS-LAST-INDEX
               ADD WS-FIRST-INDEX TO WS-LAST-INDEX GIVING WS-TEMP
               DIVIDE WS-TEMP BY 2 GIVING WS-MIDDLE-INDEX

               SET IDX TO WS-MIDDLE-INDEX
               MOVE WS-ELEMENT(IDX) TO WS-MIDDLE-ELEMENT

               IF WS-ITEM = WS-MIDDLE-ELEMENT
                   SET FOUND TO TRUE
                   MOVE WS-MIDDLE-INDEX TO WS-RESULT
                   EXIT PERFORM
               ELSE IF WS-ITEM < WS-MIDDLE-ELEMENT
                   SUBTRACT 1 FROM WS-MIDDLE-INDEX GIVING WS-LAST-INDEX
               ELSE
                   ADD 1 TO WS-MIDDLE-INDEX GIVING WS-FIRST-INDEX
               END-IF
           END-PERFORM.

           IF NOT FOUND
               MOVE "value not in array" TO WS-ERROR
           END-IF.

       DISPLAY-RESULT.
           DISPLAY "WS-FOUND-NOT-FOUND=" WS-FOUND-NOT-FOUND.
           IF FOUND
               DISPLAY "Found at position: " WS-RESULT
           ELSE
               DISPLAY WS-ERROR
           END-IF.

       PARSE-ARRAY.
           UNSTRING WS-ARRAY
               DELIMITED BY ALL ","
               INTO WS-SUBSTR
               WITH POINTER WS-PTR
               TALLYING IN WS-COUNT.

           MOVE 1 TO WS-LOOP-IDX.
           MOVE 1 TO WS-PTR.

           PERFORM VARYING WS-LOOP-IDX FROM 1 BY 1 UNTIL WS-LOOP-IDX > WS-COUNT
               UNSTRING WS-ARRAY DELIMITED BY ","
                   INTO WS-SUBSTR
                   WITH POINTER WS-PTR
               MOVE FUNCTION NUMVAL(WS-SUBSTR) TO WS-ELEMENT(WS-LOOP-IDX)
           END-PERFORM.