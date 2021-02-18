package model

import (
	"encoding/xml"
	"time"
)

// Dealing with xml:lang and xml:space, which are not recognized somehow
type Lang string
type Space string

// AboveBelow is The above-below type is used to indicate whether one element appears above or below another element.
type AboveBelow string

// BeamLevel is The MusicXML format supports six levels of beaming, up to 1024th notes. Unlike the number-level type, the beam-level type identifies concurrent beams in a beam group. It does not distinguish overlapping beams such as grace notes within regular notes, or beams used in different voices.
type BeamLevel int

// Color is The color type indicates the color of an element. Color may be represented as hexadecimal RGB triples, as in HTML, or as hexadecimal ARGB tuples, with the A indicating alpha of transparency. An alpha value of 00 is totally transparent; FF is totally opaque. If RGB is used, the A value is assumed to be FF.
//
// For instance, the RGB value "#800080" represents purple. An ARGB value of "#40800080" would be a transparent purple.
//
// As in SVG 1.1, colors are defined in terms of the sRGB color space (IEC 61966).
type Color string

// CommaSeparatedText is The comma-separated-text type is used to specify a comma-separated list of text elements, as is used by the font-family attribute.
type CommaSeparatedText string

// CssFontSize is The css-font-size type includes the CSS font sizes used as an alternative to a numeric point size.
type CssFontSize string

// Divisions is The divisions type is used to express values in terms of the musical divisions defined by the divisions element. It is preferred that these be integer values both for MIDI interoperability and to avoid roundoff errors.
type Divisions float64

// EnclosureShape is The enclosure-shape type describes the shape and presence / absence of an enclosure around text or symbols. A bracket enclosure is similar to a rectangle with the bottom line missing, as is common in jazz notation. An inverted-bracket enclosure is similar to a rectangle with the top line missing.
type EnclosureShape string

// FermataShape is The fermata-shape type represents the shape of the fermata sign. The empty value is equivalent to the normal value.
type FermataShape string

// FontSize is The font-size can be one of the CSS font sizes or a numeric point size.
type FontSize struct {
	XMLName     xml.Name `xml:"font-size"`
	Decimal     float64
	CssFontSize string
}

// FontStyle is The font-style type represents a simplified version of the CSS font-style property.
type FontStyle string

// FontWeight is The font-weight type represents a simplified version of the CSS font-weight property.
type FontWeight string

// LeftCenterRight is The left-center-right type is used to define horizontal alignment and text justification.
type LeftCenterRight string

// LeftRight is The left-right type is used to indicate whether one element appears to the left or the right of another element.
type LeftRight string

// LineLength is The line-length type distinguishes between different line lengths for doit, falloff, plop, and scoop articulations.
type LineLength string

// LineShape is The line-shape type distinguishes between straight and curved lines.
type LineShape string

// LineType is The line-type type distinguishes between solid, dashed, dotted, and wavy lines.
type LineType string

// Midi16 is The midi-16 type is used to express MIDI 1.0 values that range from 1 to 16.
type Midi16 int

// Midi128 is The midi-128 type is used to express MIDI 1.0 values that range from 1 to 128.
type Midi128 int

// Midi16384 is The midi-16384 type is used to express MIDI 1.0 values that range from 1 to 16,384.
type Midi16384 int

// Mute is The mute type represents muting for different instruments, including brass, winds, and strings. The on and off values are used for undifferentiated mutes. The remaining values represent specific mutes.
type Mute string

// NonNegativeDecimal is The non-negative-decimal type specifies a non-negative decimal value.
type NonNegativeDecimal float64

// NumberLevel is Slurs, tuplets, and many other features can be concurrent and overlapping within a single musical part. The number-level type distinguishes up to six concurrent objects of the same type. A reading program should be prepared to handle cases where the number-levels stop in an arbitrary order. Different numbers are needed when the features overlap in MusicXML document order. When a number-level value is optional, the value is 1 by default.
type NumberLevel int

// NumberOfLines is The number-of-lines type is used to specify the number of lines in text decoration attributes.
type NumberOfLines int

// NumberOrNormal is The number-or-normal values can be either a decimal number or the string "normal". This is used by the line-height and letter-spacing attributes.
type NumberOrNormal struct {
	XMLName xml.Name `xml:"number-or-normal"`
	Decimal float64
}

// OverUnder is The over-under type is used to indicate whether the tips of curved lines such as slurs and ties are overhand (tips down) or underhand (tips up).
type OverUnder string

// Percent is The percent type specifies a percentage from 0 to 100.
type Percent float64

// PositiveDecimal is The positive-decimal type specifies a positive decimal value.
type PositiveDecimal float64

// PositiveDivisions is The positive-divisions type restricts divisions values to positive numbers.
type PositiveDivisions float64

// PositiveIntegerOrEmpty is The positive-integer-or-empty values can be either a positive integer or an empty string.
type PositiveIntegerOrEmpty struct {
	XMLName         xml.Name `xml:"positive-integer-or-empty"`
	PositiveInteger int
}

// RotationDegrees is The rotation-degrees type specifies rotation, pan, and elevation values in degrees. Values range from -180 to 180.
type RotationDegrees float64

// SemiPitched is The semi-pitched type represents categories of indefinite pitch for percussion instruments.
type SemiPitched string

// SmuflGlyphName is The smufl-glyph-name type is used for attributes that reference a specific Standard Music Font Layout (SMuFL) character. The value is a SMuFL canonical glyph name, not a code point. For instance, the value for a standard piano pedal mark would be keyboardPedalPed, not U+E650.
type SmuflGlyphName string

// SmuflAccidentalGlyphName is The smufl-accidental-glyph-name type is used to reference a specific Standard Music Font Layout (SMuFL) accidental character. The value is a SMuFL canonical glyph name that starts with one of the strings used at the start of glyph names for SMuFL accidentals.
type SmuflAccidentalGlyphName string

// SmuflCodaGlyphName is The smufl-coda-glyph-name type is used to reference a specific Standard Music Font Layout (SMuFL) coda character. The value is a SMuFL canonical glyph name that starts with coda.
type SmuflCodaGlyphName string

// SmuflLyricsGlyphName is The smufl-lyrics-glyph-name type is used to reference a specific Standard Music Font Layout (SMuFL) lyrics elision character. The value is a SMuFL canonical glyph name that starts with lyrics.
type SmuflLyricsGlyphName string

// SmuflPictogramGlyphName is The smufl-pictogram-glyph-name type is used to reference a specific Standard Music Font Layout (SMuFL) percussion pictogram character. The value is a SMuFL canonical glyph name that starts with pict.
type SmuflPictogramGlyphName string

// SmuflSegnoGlyphName is The smufl-segno-glyph-name type is used to reference a specific Standard Music Font Layout (SMuFL) segno character. The value is a SMuFL canonical glyph name that starts with segno.
type SmuflSegnoGlyphName string

// StartNote is The start-note type describes the starting note of trills and mordents for playback, relative to the current note.
type StartNote string

// StartStop is The start-stop type is used for an attribute of musical elements that can either start or stop, such as tuplets.
//
// The values of start and stop refer to how an element appears in musical score order, not in MusicXML document order. An element with a stop attribute may precede the corresponding element with a start attribute within a MusicXML document. This is particularly common in multi-staff music. For example, the stopping point for a tuplet may appear in staff 1 before the starting point for the tuplet appears in staff 2 later in the document.
type StartStop string

// StartStopContinue is The start-stop-continue type is used for an attribute of musical elements that can either start or stop, but also need to refer to an intermediate point in the symbol, as for complex slurs or for formatting of symbols across system breaks.
//
// The values of start, stop, and continue refer to how an element appears in musical score order, not in MusicXML document order. An element with a stop attribute may precede the corresponding element with a start attribute within a MusicXML document. This is particularly common in multi-staff music. For example, the stopping point for a slur may appear in staff 1 before the starting point for the slur appears in staff 2 later in the document.
type StartStopContinue string

// StartStopSingle is The start-stop-single type is used for an attribute of musical elements that can be used for either multi-note or single-note musical elements, as for groupings.
type StartStopSingle string

// StringNumber is The string-number type indicates a string number. Strings are numbered from high to low, with 1 being the highest pitched full-length string.
type StringNumber int

// SymbolSize is The symbol-size type is used to distinguish between full, cue sized, grace cue sized, and oversized symbols.
type SymbolSize string

// Tenths is The tenths type is a number representing tenths of interline staff space (positive or negative). Both integer and decimal values are allowed, such as 5 for a half space and 2.5 for a quarter space. Interline space is measured from the middle of a staff line.
//
// Distances in a MusicXML file are measured in tenths of staff space. Tenths are then scaled to millimeters within the scaling element, used in the defaults element at the start of a score. Individual staves can apply a scaling factor to adjust staff size. When a MusicXML element or attribute refers to tenths, it means the global tenths defined by the scaling element, not the local tenths as adjusted by the staff-size element.
type Tenths float64

// TextDirection is The text-direction type is used to adjust and override the Unicode bidirectional text algorithm, similar to the W3C Internationalization Tag Set recommendation. Values are ltr (left-to-right embed), rtl (right-to-left embed), lro (left-to-right bidi-override), and rlo (right-to-left bidi-override). The default value is ltr. This type is typically used by applications that store text in left-to-right visual order rather than logical order. Such applications can use the lro value to better communicate with other applications that more fully support bidirectional text.
type TextDirection string

// TiedType is The tied-type type is used as an attribute of the tied element to specify where the visual representation of a tie begins and ends. A tied element which joins two notes of the same pitch can be specified with tied-type start on the first note and tied-type stop on the second note. To indicate a note should be undamped, use a single tied element with tied-type let-ring. For other ties that are visually attached to a single note, such as a tie leading into or out of a repeated section or coda, use two tied elements on the same note, one start and one stop.
//
// In start-stop cases, ties can add more elements using a continue type. This is typically used to specify the formatting of cross-system ties.
type TiedType string

// TimeOnly is The time-only type is used to indicate that a particular playback- or listening-related element only applies particular times through a repeated section. The value is a comma-separated list of positive integers arranged in ascending order, indicating which times through the repeated section that the element applies.
type TimeOnly string

// TopBottom is The top-bottom type is used to indicate the top or bottom part of a vertical shape like non-arpeggiate.
type TopBottom string

// TremoloType is The tremolo-type is used to distinguish multi-note, single-note, and unmeasured tremolos.
type TremoloType string

// TrillBeats is The trill-beats type specifies the beats used in a trill-sound or bend-sound attribute group. It is a decimal value with a minimum value of 2.
type TrillBeats float64

// TrillStep is The trill-step type describes the alternating note of trills and mordents for playback, relative to the current note.
type TrillStep string

// TwoNoteTurn is The two-note-turn type describes the ending notes of trills and mordents for playback, relative to the current note.
type TwoNoteTurn string

// UpDown is The up-down type is used for the direction of arrows and other pointed symbols like vertical accents, indicating which way the tip is pointing.
type UpDown string

// UprightInverted is The upright-inverted type describes the appearance of a fermata element. The value is upright if not specified.
type UprightInverted string

// Valign is The valign type is used to indicate vertical alignment to the top, middle, bottom, or baseline of the text. If the text is on multiple lines, baseline alignment refers to the baseline of the lowest line of text. Defaults are implementation-dependent.
type Valign string

// ValignImage is The valign-image type is used to indicate vertical alignment for images and graphics, so it does not include a baseline value. Defaults are implementation-dependent.
type ValignImage string

// YesNo is The yes-no type is used for boolean-like attributes. We cannot use W3C XML Schema booleans due to their restrictions on expression of boolean values.
type YesNo string

// YesNoNumber is The yes-no-number type is used for attributes that can be either boolean or numeric values.
type YesNoNumber struct {
	XMLName xml.Name `xml:"yes-no-number"`
	YesNo   string
	Decimal float64
}

// YyyyMmDd is Calendar dates are represented yyyy-mm-dd format, following ISO 8601. This is a W3C XML Schema date type, but without the optional timezone data.
type YyyyMmDd time.Time

// CancelLocation is The cancel-location type is used to indicate where a key signature cancellation appears relative to a new key signature: to the left, to the right, or before the barline and to the left. It is left by default. For mid-measure key elements, a cancel-location of before-barline should be treated like a cancel-location of left.
type CancelLocation string

// ClefSign is The clef-sign element represents the different clef symbols. The jianpu sign indicates that the music that follows should be in jianpu numbered notation, just as the TAB sign indicates that the music that follows should be in tablature notation. Unlike TAB, a jianpu sign does not correspond to a visual clef notation.
//
// The none sign is deprecated as of MusicXML 4.0. Use the clef element's print-object attribute instead. When the none sign is used, notes should be displayed as if in treble clef.
type ClefSign string

// Fifths is The fifths type represents the number of flats or sharps in a traditional key signature. Negative numbers are used for flats and positive numbers for sharps, reflecting the key's placement within the circle of fifths (hence the type name).
type Fifths int

// Mode is The mode type is used to specify major/minor and other mode distinctions. Valid mode values include major, minor, dorian, phrygian, lydian, mixolydian, aeolian, ionian, locrian, and none.
type Mode string

// ShowFrets is The show-frets type indicates whether to show tablature frets as numbers (0, 1, 2) or letters (a, b, c). The default choice is numbers.
type ShowFrets string

// StaffLine is The staff-line type indicates the line on a given staff. Staff lines are numbered from bottom to top, with 1 being the bottom line on a staff. Staff line values can be used to specify positions outside the staff, such as a C clef positioned in the middle of a grand staff.
type StaffLine int

// StaffNumber is The staff-number type indicates staff numbers within a multi-staff part. Staves are numbered from top to bottom, with 1 being the top staff on a part.
type StaffNumber int

// StaffType is The staff-type value can be ossia, editorial, cue, alternate, or regular. An ossia staff represents music that can be played instead of what appears on the regular staff. An editorial staff also represents musical alternatives, but is created by an editor rather than the composer. It can be used for suggested interpretations or alternatives from other sources. A cue staff represents music from another part. An alternate staff shares the same music as the prior staff, but displayed differently (e.g., treble and bass clef, standard notation and tablature). It is not included in playback. An alternate staff provides more information to an application reading a file than encoding the same music in separate parts, so its use is preferred in this situation if feasible. A regular staff is the standard default staff-stype.
type StaffType string

// TimeRelation is The time-relation type indicates the symbol used to represent the interchangeable aspect of dual time signatures.
type TimeRelation string

// TimeSeparator is The time-separator type indicates how to display the arrangement between the beats and beat-type values in a time signature. The default value is none. The horizontal, diagonal, and vertical values represent horizontal, diagonal lower-left to upper-right, and vertical lines respectively. For these values, the beats and beat-type values are arranged on either side of the separator line. The none value represents no separator with the beats and beat-type arranged vertically. The adjacent value represents no separator with the beats and beat-type arranged horizontally.
type TimeSeparator string

// TimeSymbol is The time-symbol type indicates how to display a time signature. The normal value is the usual fractional display, and is the implied symbol type if none is specified. Other options are the common and cut time symbols, as well as a single number with an implied denominator. The note symbol indicates that the beat-type should be represented with the corresponding downstem note rather than a number. The dotted-note symbol indicates that the beat-type should be represented with a dotted downstem note that corresponds to three times the beat-type value, and a numerator that is one third the beats value.
type TimeSymbol string

// BackwardForward is The backward-forward type is used to specify repeat directions. The start of the repeat has a forward direction while the end of the repeat has a backward direction.
type BackwardForward string

// BarStyle is The bar-style type represents barline style information. Choices are regular, dotted, dashed, heavy, light-light, light-heavy, heavy-light, heavy-heavy, tick (a short stroke through the top line), short (a partial barline between the 2nd and 4th lines), and none.
type BarStyle string

// EndingNumber is The ending-number type is used to specify either a comma-separated list of positive integers without leading zeros, or a string of zero or more spaces. It is used for the number attribute of the ending element. The zero or more spaces version is used when software knows that an ending is present, but cannot determine the type of the ending.
type EndingNumber string

// RightLeftMiddle is The right-left-middle type is used to specify barline location.
type RightLeftMiddle string

// StartStopDiscontinue is The start-stop-discontinue type is used to specify ending types. Typically, the start type is associated with the left barline of the first measure in an ending. The stop and discontinue types are associated with the right barline of the last measure in an ending. Stop is used when the ending mark concludes with a downward jog, as is typical for first endings. Discontinue is used when there is no downward jog, as is typical for second endings that do not conclude a piece.
type StartStopDiscontinue string

// Winged is The winged attribute indicates whether the repeat has winged extensions that appear above and below the barline. The straight and curved values represent single wings, while the double-straight and double-curved values represent double wings. The none value indicates no wings and is the default.
type Winged string

// AccordionMiddle is The accordion-middle type may have values of 1, 2, or 3, corresponding to having 1 to 3 dots in the middle section of the accordion registration symbol. This type is not used if no dots are present.
type AccordionMiddle int

// BeaterValue is The beater-value type represents pictograms for beaters, mallets, and sticks that do not have different materials represented in the pictogram. The finger and hammer values are in addition to Stone's list.
type BeaterValue string

// DegreeSymbolValue is The degree-symbol-value type indicates indicates that a symbol should be used in specifying the degree.
type DegreeSymbolValue string

// DegreeTypeValue is The degree-type-value type indicates whether the current degree element is an addition, alteration, or subtraction to the kind of the current chord in the harmony element.
type DegreeTypeValue string

// Effect is The effect type represents pictograms for sound effect percussion instruments. The cannon, lotus flute, and megaphone values are in addition to Stone's list.
type Effect string

// GlassValue is The glass-value type represents pictograms for glass percussion instruments.
type GlassValue string

// HarmonyArrangement is The harmony-arrangement type indicates how stacked chords and bass notes are displayed within a harmony element. The vertical value specifies that the second element appears below the first. The horizontal value specifies that the second element appears to the right of the first. The diagonal value specifies that the second element appears both below and to the right of the first.
type HarmonyArrangement string

// HarmonyType is The harmony-type type differentiates different types of harmonies when alternate harmonies are possible. Explicit harmonies have all note present in the music; implied have some notes missing but implied; alternate represents alternate analyses.
type HarmonyType string

// KindValue is A kind-value indicates the type of chord. Degree elements can then add, subtract, or alter from these starting points. Values include:
//
// Triads:
// major (major third, perfect fifth)
// minor (minor third, perfect fifth)
// augmented (major third, augmented fifth)
// diminished (minor third, diminished fifth)
// Sevenths:
// dominant (major triad, minor seventh)
// major-seventh (major triad, major seventh)
// minor-seventh (minor triad, minor seventh)
// diminished-seventh (diminished triad, diminished seventh)
// augmented-seventh (augmented triad, minor seventh)
// half-diminished (diminished triad, minor seventh)
// major-minor (minor triad, major seventh)
// Sixths:
// major-sixth (major triad, added sixth)
// minor-sixth (minor triad, added sixth)
// Ninths:
// dominant-ninth (dominant-seventh, major ninth)
// major-ninth (major-seventh, major ninth)
// minor-ninth (minor-seventh, major ninth)
// 11ths (usually as the basis for alteration):
// dominant-11th (dominant-ninth, perfect 11th)
// major-11th (major-ninth, perfect 11th)
// minor-11th (minor-ninth, perfect 11th)
// 13ths (usually as the basis for alteration):
// dominant-13th (dominant-11th, major 13th)
// major-13th (major-11th, major 13th)
// minor-13th (minor-11th, major 13th)
// Suspended:
// suspended-second (major second, perfect fifth)
// suspended-fourth (perfect fourth, perfect fifth)
// Functional sixths:
// Neapolitan
// Italian
// French
// German
// Other:
// pedal (pedal-point bass)
// power (perfect fifth)
// Tristan
//
// The "other" kind is used when the harmony is entirely composed of add elements.
//
// The "none" kind is used to explicitly encode absence of chords or functional harmony. In this case, the root or function element has no meaning. When using the root element, the root-step text attribute should be set to the empty string to keep the root from being displayed.
type KindValue string

// LineEnd is The line-end type specifies if there is a jog up or down (or both), an arrow, or nothing at the start or end of a bracket.
type LineEnd string

// MeasureNumberingValue is The measure-numbering-value type describes how measure numbers are displayed on this part: no numbers, numbers every measure, or numbers every system.
type MeasureNumberingValue string

// Membrane is The membrane type represents pictograms for membrane percussion instruments.
type Membrane string

// Metal is The metal type represents pictograms for metal percussion instruments. The hi-hat value refers to a pictogram like Stone's high-hat cymbals but without the long vertical line at the bottom.
type Metal string

// Milliseconds is The milliseconds type represents an integral number of milliseconds.
type Milliseconds int

// OnOff is The on-off type is used for notation elements such as string mutes.
type OnOff string

// PedalType is The pedal-type simple type is used to distinguish types of pedal directions. The start value indicates the start of a damper pedal, while the sostenuto value indicates the start of a sostenuto pedal. The other values can be used with either the damper or sostenuto pedal. The soft pedal is not included here because there is no special symbol or graphic used for it beyond what can be specified with words and bracket elements.
//
// The change, continue, discontinue, and resume types are used when the line attribute is yes. The change type indicates a pedal lift and retake indicated with an inverted V marking. The continue type allows more precise formatting across system breaks and for more complex pedaling lines. The discontinue type indicates the end of a pedal line that does not include the explicit lift represented by the stop type. The resume type indicates the start of a pedal line that does not include the downstroke represented by the start type. It can be used when a line resumes after being discontinued, or to start a pedal line that is preceded by a text or symbol representation of the pedal.
type PedalType string

// PitchedValue is The pitched-value type represents pictograms for pitched percussion instruments. The chimes and tubular chimes values distinguish the single-line and double-line versions of the pictogram.
type PitchedValue string

// PrincipalVoiceSymbol is The principal-voice-symbol type represents the type of symbol used to indicate the start of a principal or secondary voice. The "plain" value represents a plain square bracket. The value of "none" is used for analysis markup when the principal-voice element does not have a corresponding appearance in the score.
type PrincipalVoiceSymbol string

// StaffDivideSymbol is The staff-divide-symbol type is used for staff division symbols. The down, up, and up-down values correspond to SMuFL code points U+E00B, U+E00C, and U+E00D respectively.
type StaffDivideSymbol string

// StartStopChangeContinue is The start-stop-change-continue type is used to distinguish types of pedal directions.
type StartStopChangeContinue string

// SyncType is The sync-type type specifies the style that a score following application should use to synchronize an accompaniment with a performer. The none type indicates no synchronization to the performer. The tempo type indicates synchronization based on the performer tempo rather than individual events in the score. The event type indicates synchronization by following the performance of individual events in the score rather than the performer tempo. The mostly-tempo and mostly-event types combine these two approaches, with mostly-tempo giving more weight to tempo and mostly-event giving more weight to performed events. The always-event type provides the strictest synchronization by not being forgiving of missing performed events.
type SyncType string

// SystemRelationNumber is The system-relation-number type distinguishes measure numbers that are associated with a system rather than the particular part where the element appears. A value of only-top or only-bottom indicates that the number should appear only on the top or bottom part of the current system, respectively. A value of also-top or also-bottom indicates that the number should appear on both the current part and the top or bottom part of the current system, respectively. If these values appear in a score, when parts are created the number should only appear once in this part, not twice. A value of none indicates that the number is associated only with the current part, not with the system.
type SystemRelationNumber string

// SystemRelation is The system-relation type distinguishes elements that are associated with a system rather than the particular part where the element appears. A value of only-top indicates that the element should appear only on the top part of the current system. A value of also-top indicates that the element should appear on both the current part and the top part of the current system. If this value appears in a score, when parts are created the element should only appear once in this part, not twice. A value of none indicates that the element is associated only with the current part, not with the system.
type SystemRelation string

// TipDirection is The tip-direction type represents the direction in which the tip of a stick or beater points, using Unicode arrow terminology.
type TipDirection string

// StickLocation is The stick-location type represents pictograms for the location of sticks, beaters, or mallets on cymbals, gongs, drums, and other instruments.
type StickLocation string

// StickMaterial is The stick-material type represents the material being displayed in a stick pictogram.
type StickMaterial string

// StickType is The stick-type type represents the shape of pictograms where the material
// in the stick, mallet, or beater is represented in the pictogram.
type StickType string

// UpDownStopContinue is The up-down-stop-continue type is used for octave-shift elements, indicating the direction of the shift from their true pitched values because of printing difficulty.
type UpDownStopContinue string

// WedgeType is The wedge type is crescendo for the start of a wedge that is closed at the left side, diminuendo for the start of a wedge that is closed on the right side, and stop for the end of a wedge. The continue type is used for formatting wedges over a system break, or for other situations where a single wedge is divided into multiple segments.
type WedgeType string

// Wood is The wood type represents pictograms for wood percussion instruments. The maraca and maracas values distinguish the one- and two-maraca versions of the pictogram.
type Wood string

// DistanceType is The distance-type defines what type of distance is being defined in a distance element. Values include beam and hyphen. This is left as a string so that other application-specific types can be defined, but it is made a separate type so that it can be redefined more strictly.
type DistanceType string

// GlyphType is The glyph-type defines what type of glyph is being defined in a glyph element. Values include quarter-rest, g-clef-ottava-bassa, c-clef, f-clef, percussion-clef, octave-shift-up-8, octave-shift-down-8, octave-shift-continue-8, octave-shift-down-15, octave-shift-up-15, octave-shift-continue-15, octave-shift-down-22, octave-shift-up-22, and octave-shift-continue-22. This is left as a string so that other application-specific types can be defined, but it is made a separate type so that it can be redefined more strictly.
//
// A quarter-rest type specifies the glyph to use when a note has a rest element and a type value of quarter. The c-clef, f-clef, and percussion-clef types specify the glyph to use when a clef sign element value is C, F, or percussion respectively. The g-clef-ottava-bassa type specifies the glyph to use when a clef sign element value is G and the clef-octave-change element value is -1. The octave-shift types specify the glyph to use when an octave-shift type attribute value is up, down, or continue and the octave-shift size attribute value is 8, 15, or 22.
type GlyphType string

// LineWidthType is The line-width-type defines what type of line is being defined in a line-width element. Values include beam, bracket, dashes, enclosure, ending, extend, heavy barline, leger, light barline, octave shift, pedal, slur middle, slur tip, staff, stem, tie middle, tie tip, tuplet bracket, and wedge. This is left as a string so that other application-specific types can be defined, but it is made a separate type so that it can be redefined more strictly.
type LineWidthType string

// MarginType is The margin-type type specifies whether margins apply to even page, odd pages, or both.
type MarginType string

// Millimeters is The millimeters type is a number representing millimeters. This is used in the scaling element to provide a default scaling from tenths to physical units.
type Millimeters float64

// NoteSizeType is The note-size-type type indicates the type of note being defined by a note-size element. The grace-cue type is used for notes of grace-cue size. The grace type is used for notes of cue size that include a grace element. The cue type is used for all other notes with cue size, whether defined explicitly or implicitly via a cue element. The large type is used for notes of large size.
type NoteSizeType string

// AccidentalValue is The accidental-value type represents notated accidentals supported by MusicXML. In the MusicXML 2.0 DTD this was a string with values that could be included. The XSD strengthens the data typing to an enumerated list. The quarter- and three-quarters- accidentals are Tartini-style quarter-tone accidentals. The -down and -up accidentals are quarter-tone accidentals that include arrows pointing down or up. The slash- accidentals are used in Turkish classical music. The numbered sharp and flat accidentals are superscripted versions of the accidental signs, used in Turkish folk music. The sori and koron accidentals are microtonal sharp and flat accidentals used in Iranian and Persian music. The other accidental covers accidentals other than those listed here. It is usually used in combination with the smufl attribute to specify a particular SMuFL accidental. The smufl attribute may be used with any accidental value to help specify the appearance of symbols that share the same MusicXML semantics.
type AccidentalValue string

// ArrowDirection is The arrow-direction type represents the direction in which an arrow points, using Unicode arrow terminology.
type ArrowDirection string

// ArrowStyle is The arrow-style type represents the style of an arrow, using Unicode arrow terminology. Filled and hollow arrows indicate polygonal single arrows. Paired arrows are duplicate single arrows in the same direction. Combined arrows apply to double direction arrows like left right, indicating that an arrow in one direction should be combined with an arrow in the other direction.
type ArrowStyle string

// BeamValue is The beam-value type represents the type of beam associated with each of 8 beam levels (up to 1024th notes) available for each note.
type BeamValue string

// BendShape is The bend-shape type distinguishes between the angled bend symbols commonly used in standard notation and the curved bend symbols commonly used in both tablature and standard notation.
type BendShape string

// BreathMarkValue is The breath-mark-value type represents the symbol used for a breath mark.
type BreathMarkValue string

// CaesuraValue is The caesura-value type represents the shape of the caesura sign.
type CaesuraValue string

// CircularArrow is The circular-arrow type represents the direction in which a circular arrow points, using Unicode arrow terminology.
type CircularArrow string

// Fan is The fan type represents the type of beam fanning present on a note, used to represent accelerandos and ritardandos.
type Fan string

// HandbellValue is The handbell-value type represents the type of handbell technique being notated.
type HandbellValue string

// HarmonClosedLocation is The harmon-closed-location type indicates which portion of the symbol is filled in when the corresponding harmon-closed-value is half.
type HarmonClosedLocation string

// HarmonClosedValue is The harmon-closed-value type represents whether the harmon mute is closed, open, or half-open.
type HarmonClosedValue string

// HoleClosedLocation is The hole-closed-location type indicates which portion of the hole is filled in when the corresponding hole-closed-value is half.
type HoleClosedLocation string

// HoleClosedValue is The hole-closed-value type represents whether the hole is closed, open, or half-open.
type HoleClosedValue string

// NoteTypeValue is The note-type-value type is used for the MusicXML type element and represents the graphic note type, from 1024th (shortest) to maxima (longest).
type NoteTypeValue string

// NoteheadValue is The notehead-value type indicates shapes other than the open and closed ovals associated with note durations.
//
// The values do, re, mi, fa, fa up, so, la, and ti correspond to Aikin's 7-shape system.  The fa up shape is typically used with upstems; the fa shape is typically used with downstems or no stems.
//
// The arrow shapes differ from triangle and inverted triangle by being centered on the stem. Slashed and back slashed notes include both the normal notehead and a slash. The triangle shape has the tip of the triangle pointing up; the inverted triangle shape has the tip of the triangle pointing down. The left triangle shape is a right triangle with the hypotenuse facing up and to the left.
//
// The other notehead covers noteheads other than those listed here. It is usually used in combination with the smufl attribute to specify a particular SMuFL notehead. The smufl attribute may be used with any notehead value to help specify the appearance of symbols that share the same MusicXML semantics. Noteheads in the SMuFL "Note name noteheads" range (U+E150–U+E1AF) should not use the smufl attribute or the "other" value, but instead use the notehead-text element.
type NoteheadValue string

// Octave is Octaves are represented by the numbers 0 to 9, where 4 indicates the octave started by middle C.
type Octave int

// Semitones is The semitones type is a number representing semitones, used for chromatic alteration. A value of -1 corresponds to a flat and a value of 1 to a sharp. Decimal values like 0.5 (quarter tone sharp) are used for microtones.
type Semitones float64

// ShowTuplet is The show-tuplet type indicates whether to show a part of a tuplet relating to the tuplet-actual element, both the tuplet-actual and tuplet-normal elements, or neither.
type ShowTuplet string

// StemValue is The stem type represents the notated stem direction.
type StemValue string

// Step is The step type represents a step of the diatonic scale, represented using the English letters A through G.
type Step string

// Syllabic is Lyric hyphenation is indicated by the syllabic type. The single, begin, end, and middle values represent single-syllable words, word-beginning syllables, word-ending syllables, and mid-word syllables, respectively.
type Syllabic string

// TapHand is The tap-hand type represents the symbol to use for a tap element. The left and right values refer to the SMuFL guitarLeftHandTapping and guitarRightHandTapping glyphs respectively.
type TapHand string

// TremoloMarks is The number of tremolo marks is represented by a number from 0 to 8: the same as beam-level with 0 added.
type TremoloMarks int

// GroupBarlineValue is The group-barline-value type indicates if the group should have common barlines.
type GroupBarlineValue string

// GroupSymbolValue is The group-symbol-value type indicates how the symbol for a group is indicated in the score. The default value is none.
type GroupSymbolValue string

// MeasureText is The measure-text type is used for the text attribute of measure elements. It has at least one character. The implicit attribute of the measure element should be set to "yes" rather than setting the text attribute to an empty string.
type MeasureText string

// BendSound is The bend-sound type is used for bend and slide elements, and is similar to the trill-sound attribute group. Here the beats element refers to the number of discrete elements (like MIDI pitch bends) used to represent a continuous bend or slide. The first-beat indicates the percentage of the direction for starting a bend; the last-beat the percentage for ending it. The default choices are:
//
// accelerate = "no"
// beats = "4"
// first-beat = "25"
// last-beat = "75"
type BendSound struct {
	XMLName        xml.Name `xml:"bend-sound"`
	AccelerateAttr string   `xml:"accelerate,attr,omitempty"`
	BeatsAttr      float64  `xml:"beats,attr,omitempty"`
	FirstBeatAttr  float64  `xml:"first-beat,attr,omitempty"`
	LastBeatAttr   float64  `xml:"last-beat,attr,omitempty"`
}

// Bezier is The bezier attribute group is used to indicate the curvature of slurs and ties, representing the control points for a cubic bezier curve. For ties, the bezier attribute group is used with the tied element.
//
// Normal slurs, S-shaped slurs, and ties need only two bezier points: one associated with the start of the slur or tie, the other with the stop. Complex slurs and slurs divided over system breaks can specify additional bezier data at slur elements with a continue type.
//
// The bezier-x, bezier-y, and bezier-offset attributes describe the outgoing bezier point for slurs and ties with a start type, and the incoming bezier point for slurs and ties with types of stop or continue. The bezier-x2, bezier-y2, and bezier-offset2 attributes are only valid with slurs of type continue, and describe the outgoing bezier point.
//
// The bezier-x, bezier-y, bezier-x2, and bezier-y2 attributes are specified in tenths, relative to any position settings associated with the slur or tied element. The bezier-offset and bezier-offset2 attributes are measured in terms of musical divisions, like the offset element.
//
// The bezier-offset and bezier-offset2 attributes are deprecated as of MusicXML 3.1. If both the bezier-x and bezier-offset attributes are present, the bezier-x attribute takes priority. Similarly, the bezier-x2 attribute takes priority over the bezier-offset2 attribute. The two types of bezier attributes are not additive.
type Bezier struct {
	XMLName           xml.Name `xml:"bezier"`
	BezierXAttr       float64  `xml:"bezier-x,attr,omitempty"`
	BezierYAttr       float64  `xml:"bezier-y,attr,omitempty"`
	BezierX2Attr      float64  `xml:"bezier-x2,attr,omitempty"`
	BezierY2Attr      float64  `xml:"bezier-y2,attr,omitempty"`
	BezierOffsetAttr  float64  `xml:"bezier-offset,attr,omitempty"`
	BezierOffset2Attr float64  `xml:"bezier-offset2,attr,omitempty"`
}

// DashedFormatting is The dashed-formatting entity represents the length of dashes and spaces in a dashed line. Both the dash-length and space-length attributes are represented in tenths. These attributes are ignored if the corresponding line-type attribute is not dashed.
type DashedFormatting struct {
	XMLName         xml.Name `xml:"dashed-formatting"`
	DashLengthAttr  float64  `xml:"dash-length,attr,omitempty"`
	SpaceLengthAttr float64  `xml:"space-length,attr,omitempty"`
}

// Directive is The directive attribute changes the default-x position of a direction. It indicates that the left-hand side of the direction is aligned with the left-hand side of the time signature. If no time signature is present, it is aligned with the left-hand side of the first music notational element in the measure. If a default-x, justify, or halign attribute is present, it overrides the directive attribute.
type Directive struct {
	XMLName       xml.Name `xml:"directive"`
	DirectiveAttr string   `xml:"directive,attr,omitempty"`
}

// DocumentAttributes is The document-attributes attribute group is used to specify the attributes for an entire MusicXML document. Currently this is used for the version attribute.
//
// The version attribute was added in Version 1.1 for the score-partwise and score-timewise documents. It provides an easier way to get version information than through the MusicXML public ID. The default value is 1.0 to make it possible for programs that handle later versions to distinguish earlier version files reliably. Programs that write MusicXML 1.1 or later files should set this attribute.
type DocumentAttributes struct {
	XMLName     xml.Name `xml:"document-attributes"`
	VersionAttr string   `xml:"version,attr,omitempty"`
}

// Enclosure is The enclosure attribute group is used to specify the formatting of an enclosure around text or symbols.
type Enclosure struct {
	XMLName       xml.Name `xml:"enclosure"`
	EnclosureAttr string   `xml:"enclosure,attr,omitempty"`
}

// Font is The font attribute group gathers together attributes for determining the font within a credit or direction. They are based on the text styles for Cascading Style Sheets. The font-family is a comma-separated list of font names. These can be specific font styles such as Maestro or Opus, or one of several generic font styles: music, engraved, handwritten, text, serif, sans-serif, handwritten, cursive, fantasy, and monospace. The music, engraved, and handwritten values refer to music fonts; the rest refer to text fonts. The fantasy style refers to decorative text such as found in older German-style printing. The font-style can be normal or italic. The font-size can be one of the CSS sizes (xx-small, x-small, small, medium, large, x-large, xx-large) or a numeric point size. The font-weight can be normal or bold. The default is application-dependent, but is a text font vs. a music font.
type Font struct {
	XMLName        xml.Name  `xml:"font"`
	FontFamilyAttr string    `xml:"font-family,attr,omitempty"`
	FontStyleAttr  string    `xml:"font-style,attr,omitempty"`
	FontSizeAttr   *FontSize `xml:"font-size,attr,omitempty"`
	FontWeightAttr string    `xml:"font-weight,attr,omitempty"`
}

// Halign is In cases where text extends over more than one line, horizontal alignment and justify values can be different. The most typical case is for credits, such as:
//
// Words and music by
//   Pat Songwriter
//
// Typically this type of credit is aligned to the right, so that the position information refers to the right-most part of the text. But in this example, the text is center-justified, not right-justified.
//
// The halign attribute is used in these situations. If it is not present, its value is the same as for the justify attribute. For elements where a justify attribute is not allowed, the default is implementation-dependent.
type Halign struct {
	XMLName    xml.Name `xml:"halign"`
	HalignAttr string   `xml:"halign,attr,omitempty"`
}

// Justify is The justify attribute is used to indicate left, center, or right justification. The default value varies for different elements. For elements where the justify attribute is present but the halign attribute is not, the justify attribute indicates horizontal alignment as well as justification.
type Justify struct {
	XMLName     xml.Name `xml:"justify"`
	JustifyAttr string   `xml:"justify,attr,omitempty"`
}

// LetterSpacing is The letter-spacing attribute specifies text tracking. Values are either "normal" or a number representing the number of ems to add between each letter. The number may be negative in order to subtract space. The default is normal, which allows flexibility of letter-spacing for purposes of text justification.
type LetterSpacing struct {
	XMLName           xml.Name        `xml:"letter-spacing"`
	LetterSpacingAttr *NumberOrNormal `xml:"letter-spacing,attr,omitempty"`
}

// LevelDisplay is The level-display attribute group specifies three common ways to indicate editorial indications: putting parentheses or square brackets around a symbol, or making the symbol a different size. If not specified, they are left to application defaults. It is used by the level and accidental elements.
type LevelDisplay struct {
	XMLName         xml.Name `xml:"level-display"`
	ParenthesesAttr string   `xml:"parentheses,attr,omitempty"`
	BracketAttr     string   `xml:"bracket,attr,omitempty"`
	SizeAttr        string   `xml:"size,attr,omitempty"`
}

// LineHeight is The line-height attribute specifies text leading. Values are either "normal" or a number representing the percentage of the current font height to use for leading. The default is "normal". The exact normal value is implementation-dependent, but values between 100 and 120 are recommended.
type LineHeight struct {
	XMLName        xml.Name        `xml:"line-height"`
	LineHeightAttr *NumberOrNormal `xml:"line-height,attr,omitempty"`
}

// OptionalUniqueId is The optional-unique-id attribute group allows an element to optionally specify an ID that is unique to the entire document. This attribute group is not used for a required id attribute, or for an id attribute that specifies an id reference.
type OptionalUniqueId struct {
	XMLName xml.Name `xml:"optional-unique-id"`
	IdAttr  string   `xml:"id,attr,omitempty"`
}

// Orientation is The orientation attribute indicates whether slurs and ties are overhand (tips down) or underhand (tips up). This is distinct from the placement attribute used by any notation type.
type Orientation struct {
	XMLName         xml.Name `xml:"orientation"`
	OrientationAttr string   `xml:"orientation,attr,omitempty"`
}

// Placement is The placement attribute indicates whether something is above or below another element, such as a note or a notation.
type Placement struct {
	XMLName       xml.Name `xml:"placement"`
	PlacementAttr string   `xml:"placement,attr,omitempty"`
}

// Position is The position attributes are based on MuseData print suggestions. For most elements, any program will compute a default x and y position. The position attributes let this be changed two ways.
//
// The default-x and default-y attributes change the computation of the default position. For most elements, the origin is changed relative to the left-hand side of the note or the musical position within the bar (x) and the top line of the staff (y).
//
// For the following elements, the default-x value changes the origin relative to the start of the current measure:
//
// - note
// - figured-bass
// - harmony
// - link
// - directive
// - measure-numbering
// - all descendants of the part-list element
// - all children of the direction-type element
//
// This origin is from the start of the entire measure, at either the left barline or the start of the system.
//
// When the default-x attribute is used within a child element of the part-name-display, part-abbreviation-display, group-name-display, or group-abbreviation-display elements, it changes the origin relative to the start of the first measure on the system. These values are used when the current measure or a succeeding measure starts a new system. The same change of origin is used for the group-symbol element.
//
// For the note, figured-bass, and harmony elements, the default-x value is considered to have adjusted the musical position within the bar for its descendant elements.
//
// Since the credit-words and credit-image elements are not related to a measure, in these cases the default-x and default-y attributes adjust the origin relative to the bottom left-hand corner of the specified page.
//
// The relative-x and relative-y attributes change the position relative to the default position, either as computed by the individual program, or as overridden by the default-x and default-y attributes.
//
// Positive x is right, negative x is left; positive y is up, negative y is down. All units are in tenths of interline space. For stems, positive relative-y lengthens a stem while negative relative-y shortens it.
//
// The default-x and default-y position attributes provide higher-resolution positioning data than related features such as the placement attribute and the offset element. Applications reading a MusicXML file that can understand both features should generally rely on the default-x and default-y attributes for their greater accuracy. For the relative-x and relative-y attributes, the offset element, placement attribute, and directive attribute provide context for the relative position information, so the two features should be interpreted together.
//
// As elsewhere in the MusicXML format, tenths are the global tenths defined by the scaling element, not the local tenths of a staff resized by the staff-size element.
type Position struct {
	XMLName       xml.Name `xml:"position"`
	DefaultXAttr  float64  `xml:"default-x,attr,omitempty"`
	DefaultYAttr  float64  `xml:"default-y,attr,omitempty"`
	RelativeXAttr float64  `xml:"relative-x,attr,omitempty"`
	RelativeYAttr float64  `xml:"relative-y,attr,omitempty"`
}

// PrintObject is The print-object attribute specifies whether or not to print an object (e.g. a note or a rest). It is yes by default.
type PrintObject struct {
	XMLName         xml.Name `xml:"print-object"`
	PrintObjectAttr string   `xml:"print-object,attr,omitempty"`
}

// PrintSpacing is The print-spacing attribute controls whether or not spacing is left for an invisible note or object. It is used only if no note, dot, or lyric is being printed. The value is yes (leave spacing) by default.
type PrintSpacing struct {
	XMLName          xml.Name `xml:"print-spacing"`
	PrintSpacingAttr string   `xml:"print-spacing,attr,omitempty"`
}

// PrintStyle is The print-style attribute group collects the most popular combination of printing attributes: position, font, and color.
type PrintStyle struct {
	XMLName xml.Name `xml:"print-style"`
}

// PrintStyleAlign is The print-style-align attribute group adds the halign and valign attributes to the position, font, and color attributes.
type PrintStyleAlign struct {
	XMLName xml.Name `xml:"print-style-align"`
}

// Printout is The printout attribute group collects the different controls over printing an object (e.g. a note or rest) and its parts, including augmentation dots and lyrics. This is especially useful for notes that overlap in different voices, or for chord sheets that contain lyrics and chords but no melody.
//
// By default, all these attributes are set to yes. If print-object is set to no, the print-dot and print-lyric attributes are interpreted to also be set to no if they are not present.
type Printout struct {
	XMLName        xml.Name `xml:"printout"`
	PrintDotAttr   string   `xml:"print-dot,attr,omitempty"`
	PrintLyricAttr string   `xml:"print-lyric,attr,omitempty"`
}

// Smufl is The smufl attribute group is used to indicate a particular Standard Music Font Layout (SMuFL) character. Sometimes this is a formatting choice, and sometimes this is a refinement of the semantic meaning of an element.
type Smufl struct {
	XMLName   xml.Name `xml:"smufl"`
	SmuflAttr string   `xml:"smufl,attr,omitempty"`
}

// SwingTypeValue is The swing-type-value attribute group specifies the note type, either eighth or 16th, to which the ratio defined in the swing element is applied.
type SwingTypeValue string

// TextDecoration ...
type TextDecoration struct {
	XMLName xml.Name `xml:"text-decoration"`
}

// TextRotation ...
type TextRotation struct {
	XMLName xml.Name `xml:"text-rotation"`
}

// SymbolFormatting is The symbol-formatting attribute group collects the common formatting attributes for musical symbols. Default values may differ across the elements that use this group.
type SymbolFormatting struct {
	XMLName xml.Name `xml:"symbol-formatting"`
}

// TextFormatting is The text-formatting attribute group collects the common formatting attributes for text elements. Default values may differ across the elements that use this group.
type TextFormatting struct {
	XMLName      xml.Name `xml:"text-formatting"`
	XmlLangAttr  *Lang    `xml:"xml:lang,attr,omitempty"`
	XmlSpaceAttr *Space   `xml:"xml:space,attr,omitempty"`
}

// TrillSound is The trill-sound attribute group includes attributes used to guide the sound of trills, mordents, turns, shakes, and wavy lines, based on MuseData sound suggestions. The default choices are:
//
// start-note = "upper"
// trill-step = "whole"
// two-note-turn = "none"
// accelerate = "no"
// beats = "4".
//
// Second-beat and last-beat are percentages for landing on the indicated beat, with defaults of 25 and 75 respectively.
//
// For mordent and inverted-mordent elements, the defaults are different:
//
// The default start-note is "main", not "upper".
// The default for beats is "3", not "4".
// The default for second-beat is "12", not "25".
// The default for last-beat is "24", not "75".
type TrillSound struct {
	XMLName         xml.Name `xml:"trill-sound"`
	StartNoteAttr   string   `xml:"start-note,attr,omitempty"`
	TrillStepAttr   string   `xml:"trill-step,attr,omitempty"`
	TwoNoteTurnAttr string   `xml:"two-note-turn,attr,omitempty"`
	AccelerateAttr  string   `xml:"accelerate,attr,omitempty"`
	BeatsAttr       float64  `xml:"beats,attr,omitempty"`
	SecondBeatAttr  float64  `xml:"second-beat,attr,omitempty"`
	LastBeatAttr    float64  `xml:"last-beat,attr,omitempty"`
}

// XPosition is The x-position attribute group is used for elements like notes where specifying x position is common, but specifying y position is rare.
type XPosition struct {
	XMLName       xml.Name `xml:"x-position"`
	DefaultXAttr  float64  `xml:"default-x,attr,omitempty"`
	DefaultYAttr  float64  `xml:"default-y,attr,omitempty"`
	RelativeXAttr float64  `xml:"relative-x,attr,omitempty"`
	RelativeYAttr float64  `xml:"relative-y,attr,omitempty"`
}

// YPosition is The y-position attribute group is used for elements like stems where specifying y position is common, but specifying x position is rare.
type YPosition struct {
	XMLName       xml.Name `xml:"y-position"`
	DefaultXAttr  float64  `xml:"default-x,attr,omitempty"`
	DefaultYAttr  float64  `xml:"default-y,attr,omitempty"`
	RelativeXAttr float64  `xml:"relative-x,attr,omitempty"`
	RelativeYAttr float64  `xml:"relative-y,attr,omitempty"`
}

// ImageAttributes is The image-attributes group is used to include graphical images in a score. The required source attribute is the URL for the image file. The required type attribute is the MIME type for the image file format. Typical choices include application/postscript, image/gif, image/jpeg, image/png, and image/tiff. The optional height and width attributes are used to size and scale an image. The image should be scaled independently in X and Y if both height and width are specified. If only one attribute is specified, the image should be scaled proportionally to fit in the specified dimension.
type ImageAttributes struct {
	XMLName    xml.Name `xml:"image-attributes"`
	SourceAttr string   `xml:"source,attr"`
	TypeAttr   string   `xml:"type,attr"`
	HeightAttr float64  `xml:"height,attr,omitempty"`
	WidthAttr  float64  `xml:"width,attr,omitempty"`
}

// PrintAttributes is The print-attributes group is used by the print element. The new-system and new-page attributes indicate whether to force a system or page break, or to force the current music onto the same system or page as the preceding music. Normally this is the first music data within a measure. If used in multi-part music, they should be placed in the same positions within each part, or the results are undefined. The page-number attribute sets the number of a new page; it is ignored if new-page is not "yes". Version 2.0 adds a blank-page attribute. This is a positive integer value that specifies the number of blank pages to insert before the current measure. It is ignored if new-page is not "yes". These blank pages have no music, but may have text or images specified by the credit element. This is used to allow a combination of pages that are all text, or all text and images, together with pages of music.
//
// The staff-spacing attribute specifies spacing between multiple staves in tenths of staff space. This is deprecated as of Version 1.1; the staff-layout element should be used instead. If both are present, the staff-layout values take priority.
type PrintAttributes struct {
	XMLName          xml.Name `xml:"print-attributes"`
	StaffSpacingAttr float64  `xml:"staff-spacing,attr,omitempty"`
	NewSystemAttr    string   `xml:"new-system,attr,omitempty"`
	NewPageAttr      string   `xml:"new-page,attr,omitempty"`
	BlankPageAttr    int      `xml:"blank-page,attr,omitempty"`
	PageNumberAttr   string   `xml:"page-number,attr,omitempty"`
}

// ElementPosition is The element and position attributes are new as of Version 2.0. They allow for bookmarks and links to be positioned at higher resolution than the level of music-data elements. When no element and position attributes are present, the bookmark or link element refers to the next sibling element in the MusicXML file. The element attribute specifies an element type for a descendant of the next sibling element that is not a link or bookmark. The position attribute specifies the position of this descendant element, where the first position is 1. The position attribute is ignored if the element attribute is not present. For instance, an element value of "beam" and a position value of "2" defines the link or bookmark to refer to the second beam descendant of the next sibling element that is not a link or bookmark. This is equivalent to an XPath test of [.//beam[2]] done in the context of the sibling element.
type ElementPosition struct {
	XMLName      xml.Name `xml:"element-position"`
	ElementAttr  string   `xml:"element,attr,omitempty"`
	PositionAttr int      `xml:"position,attr,omitempty"`
}

// LinkAttributes is The link-attributes group includes all the simple XLink attributes supported in the MusicXML format. It is also used to connect a MusicXML score with MusicXML parts or a MusicXML opus.
type LinkAttributes struct {
	XMLName          xml.Name `xml:"link-attributes"`
	XlinkHrefAttr    string   `xml:"xlink:href,attr"`
	XlinkTypeAttr    string   `xml:"xlink:type,attr,omitempty"`
	XlinkRoleAttr    string   `xml:"xlink:role,attr,omitempty"`
	XlinkTitleAttr   string   `xml:"xlink:title,attr,omitempty"`
	XlinkShowAttr    string   `xml:"xlink:show,attr,omitempty"`
	XlinkActuateAttr string   `xml:"xlink:actuate,attr,omitempty"`
}

// GroupNameText is The group-name-text attribute group is used by the group-name and group-abbreviation elements. The print-style and justify attribute groups are deprecated in MusicXML 2.0 in favor of the new group-name-display and group-abbreviation-display elements.
type GroupNameText struct {
	XMLName xml.Name `xml:"group-name-text"`
}

// MeasureAttributes is The measure-attributes group is used by the measure element. Measures have a required number attribute (going from partwise to timewise, measures are grouped via the number).
//
// The implicit attribute is set to "yes" for measures where the measure number should never appear, such as pickup measures and the last half of mid-measure repeats. The value is "no" if not specified.
//
// The non-controlling attribute is intended for use in multimetric music like the Don Giovanni minuet. If set to "yes", the left barline in this measure does not coincide with the left barline of measures in other parts. The value is "no" if not specified.
//
// In partwise files, the number attribute should be the same for measures in different parts that share the same left barline. While the number attribute is often numeric, it does not have to be. Non-numeric values are typically used together with the implicit or non-controlling attributes being set to "yes". For a pickup measure, the number attribute is typically set to "0" and the implicit attribute is typically set to "yes".
//
// If measure numbers are not unique within a part, this can cause problems for conversions between partwise and timewise formats. The text attribute allows specification of displayed measure numbers that are different than what is used in the number attribute. This attribute is ignored for measures where the implicit attribute is set to "yes". Further details about measure numbering can be specified using the measure-numbering element.
//
// Measure width is specified in tenths. These are the global tenths specified in the scaling element, not local tenths as modified by the staff-size element.The width covers the entire measure from barline or system start to barline or system end.
type MeasureAttributes struct {
	XMLName            xml.Name `xml:"measure-attributes"`
	NumberAttr         string   `xml:"number,attr"`
	TextAttr           string   `xml:"text,attr,omitempty"`
	ImplicitAttr       string   `xml:"implicit,attr,omitempty"`
	NonControllingAttr string   `xml:"non-controlling,attr,omitempty"`
	WidthAttr          float64  `xml:"width,attr,omitempty"`
}

// PartAttributes is In either partwise or timewise format, the part element has an id attribute that is an IDREF back to a score-part in the part-list.
type PartAttributes struct {
	XMLName xml.Name `xml:"part-attributes"`
	IdAttr  string   `xml:"id,attr"`
}

// PartNameText is The part-name-text attribute group is used by the part-name and part-abbreviation elements. The print-style and justify attribute groups are deprecated in MusicXML 2.0 in favor of the new part-name-display and part-abbreviation-display elements.
type PartNameText struct {
	XMLName xml.Name `xml:"part-name-text"`
}

// AccidentalText is The accidental-text type represents an element with an accidental value and text-formatting attributes.
type AccidentalText struct {
	XMLName        xml.Name `xml:"accidental-text"`
	TextFormatting *TextFormatting
	SmuflAttr      string `xml:"smufl,attr,omitempty"`
}

// Coda is The coda type is the visual indicator of a coda sign. The exact glyph can be specified with the smufl attribute. A sound element is also needed to guide playback applications reliably.
type Coda struct {
	XMLName          xml.Name `xml:"coda"`
	PrintStyleAlign  *PrintStyleAlign
	OptionalUniqueId *OptionalUniqueId
	SmuflAttr        string `xml:"smufl,attr,omitempty"`
}

// Dynamics is Dynamics can be associated either with a note or a general musical direction. To avoid inconsistencies between and amongst the letter abbreviations for dynamics (what is sf vs. sfz, standing alone or with a trailing dynamic that is not always piano), we use the actual letters as the names of these dynamic elements. The other-dynamics element allows other dynamic marks that are not covered here, but many of those should perhaps be included in a more general musical direction element. Dynamics elements may also be combined to create marks not covered by a single element, such as sfmp.
//
// These letter dynamic symbols are separated from crescendo, decrescendo, and wedge indications. Dynamic representation is inconsistent in scores. Many things are assumed by the composer and left out, such as returns to original dynamics. Systematic representations are quite complex: for example, Humdrum has at least 3 representation formats related to dynamics. The MusicXML format captures what is in the score, but does not try to be optimal for analysis or synthesis of dynamics.
//
// The placement attribute is used when the dynamics are associated with a note. It is ignored when the dynamics are associated with a direction. In that case the direction element's placement attribute is used instead.
type Dynamics struct {
	XMLName          xml.Name `xml:"dynamics"`
	PrintStyleAlign  *PrintStyleAlign
	Placement        *Placement
	TextDecoration   *TextDecoration
	Enclosure        *Enclosure
	OptionalUniqueId *OptionalUniqueId
	P                *Empty     `xml:"p"`
	Pp               *Empty     `xml:"pp"`
	Ppp              *Empty     `xml:"ppp"`
	Pppp             *Empty     `xml:"pppp"`
	Ppppp            *Empty     `xml:"ppppp"`
	Pppppp           *Empty     `xml:"pppppp"`
	F                *Empty     `xml:"f"`
	Ff               *Empty     `xml:"ff"`
	Fff              *Empty     `xml:"fff"`
	Ffff             *Empty     `xml:"ffff"`
	Fffff            *Empty     `xml:"fffff"`
	Ffffff           *Empty     `xml:"ffffff"`
	Mp               *Empty     `xml:"mp"`
	Mf               *Empty     `xml:"mf"`
	Sf               *Empty     `xml:"sf"`
	Sfp              *Empty     `xml:"sfp"`
	Sfpp             *Empty     `xml:"sfpp"`
	Fp               *Empty     `xml:"fp"`
	Rf               *Empty     `xml:"rf"`
	Rfz              *Empty     `xml:"rfz"`
	Sfz              *Empty     `xml:"sfz"`
	Sffz             *Empty     `xml:"sffz"`
	Fz               *Empty     `xml:"fz"`
	N                *Empty     `xml:"n"`
	Pf               *Empty     `xml:"pf"`
	Sfzp             *Empty     `xml:"sfzp"`
	OtherDynamics    *OtherText `xml:"other-dynamics"`
}

// Empty is The empty type represents an empty element with no attributes.
type Empty struct {
	XMLName xml.Name `xml:"empty"`
}

// EmptyPlacement is The empty-placement type represents an empty element with print-style and placement attributes.
type EmptyPlacement struct {
	XMLName    xml.Name `xml:"empty-placement"`
	PrintStyle *PrintStyle
	Placement  *Placement
}

// EmptyPlacementSmufl is The empty-placement-smufl type represents an empty element with print-style, placement, and smufl attributes.
type EmptyPlacementSmufl struct {
	XMLName    xml.Name `xml:"empty-placement-smufl"`
	PrintStyle *PrintStyle
	Placement  *Placement
	Smufl      *Smufl
}

// EmptyPrintStyle is The empty-print-style type represents an empty element with print-style attribute group.
type EmptyPrintStyle struct {
	XMLName    xml.Name `xml:"empty-print-style"`
	PrintStyle *PrintStyle
}

// EmptyPrintStyleAlign is The empty-print-style-align type represents an empty element with print-style-align attribute group.
type EmptyPrintStyleAlign struct {
	XMLName         xml.Name `xml:"empty-print-style-align"`
	PrintStyleAlign *PrintStyleAlign
}

// EmptyPrintStyleAlignId is The empty-print-style-align-id type represents an empty element with print-style-align and optional-unique-id attribute groups.
type EmptyPrintStyleAlignId struct {
	XMLName          xml.Name `xml:"empty-print-style-align-id"`
	PrintStyleAlign  *PrintStyleAlign
	OptionalUniqueId *OptionalUniqueId
}

// EmptyPrintObjectStyleAlign is The empty-print-style-align-object type represents an empty element with print-object and print-style-align attribute groups.
type EmptyPrintObjectStyleAlign struct {
	XMLName         xml.Name `xml:"empty-print-object-style-align"`
	PrintObject     *PrintObject
	PrintStyleAlign *PrintStyleAlign
}

// EmptyTrillSound is The empty-trill-sound type represents an empty element with print-style, placement, and trill-sound attributes.
type EmptyTrillSound struct {
	XMLName    xml.Name `xml:"empty-trill-sound"`
	PrintStyle *PrintStyle
	Placement  *Placement
	TrillSound *TrillSound
}

// HorizontalTurn is The horizontal-turn type represents turn elements that are horizontal rather than vertical. These are empty elements with print-style, placement, trill-sound, and slash attributes. If the slash attribute is yes, then a vertical line is used to slash the turn; it is no by default.
type HorizontalTurn struct {
	XMLName    xml.Name `xml:"horizontal-turn"`
	PrintStyle *PrintStyle
	Placement  *Placement
	TrillSound *TrillSound
	SlashAttr  string `xml:"slash,attr,omitempty"`
}

// Fermata is The fermata text content represents the shape of the fermata sign. An empty fermata element represents a normal fermata. The fermata type is upright if not specified.
type Fermata struct {
	XMLName          xml.Name `xml:"fermata"`
	PrintStyle       *PrintStyle
	OptionalUniqueId *OptionalUniqueId
	TypeAttr         string `xml:"type,attr,omitempty"`
}

// Fingering is Fingering is typically indicated 1,2,3,4,5. Multiple fingerings may be given, typically to substitute fingerings in the middle of a note. The substitution and alternate values are "no" if the attribute is not present. For guitar and other fretted instruments, the fingering element represents the fretting finger; the pluck element represents the plucking finger.
type Fingering struct {
	XMLName          xml.Name `xml:"fingering"`
	PrintStyle       *PrintStyle
	Placement        *Placement
	SubstitutionAttr string `xml:"substitution,attr,omitempty"`
	AlternateAttr    string `xml:"alternate,attr,omitempty"`
}

// FormattedSymbol is The formatted-symbol type represents a SMuFL musical symbol element with formatting attributes.
type FormattedSymbol struct {
	XMLName          xml.Name `xml:"formatted-symbol"`
	SymbolFormatting *SymbolFormatting
}

// FormattedSymbolId is The formatted-symbol-id type represents a SMuFL musical symbol element with formatting and id attributes.
type FormattedSymbolId struct {
	XMLName          xml.Name `xml:"formatted-symbol-id"`
	SymbolFormatting *SymbolFormatting
	OptionalUniqueId *OptionalUniqueId
}

// FormattedText is The formatted-text type represents a text element with text-formatting attributes.
type FormattedText struct {
	XMLName        xml.Name `xml:"formatted-text"`
	TextFormatting *TextFormatting
}

// FormattedTextId is The formatted-text-id type represents a text element with text-formatting and id attributes.
type FormattedTextId struct {
	XMLName          xml.Name `xml:"formatted-text-id"`
	TextFormatting   *TextFormatting
	OptionalUniqueId *OptionalUniqueId
}

// Fret is The fret element is used with tablature notation and chord diagrams. Fret numbers start with 0 for an open string and 1 for the first fret.
type Fret struct {
	XMLName xml.Name `xml:"fret"`
	Font    *Font
	Color   string
}

// Level is The level type is used to specify editorial information for different MusicXML elements. The content contains identifying and/or descriptive text about the editorial status of the parent element.
//
// If the reference attribute is yes, this indicates editorial information that is for display only and should not affect playback. For instance, a modern edition of older music may set reference="yes" on the attributes containing the music's original clef, key, and time signature. It is no by default.
//
// The type attribute indicates whether the editorial information applies to the start of a series of symbols, the end of a series of symbols, or a single symbol. It is single by default for compatibility with earlier MusicXML versions.
type Level struct {
	XMLName       xml.Name `xml:"level"`
	LevelDisplay  *LevelDisplay
	ReferenceAttr string `xml:"reference,attr,omitempty"`
	TypeAttr      string `xml:"type,attr,omitempty"`
}

// MidiDevice is The midi-device type corresponds to the DeviceName meta event in Standard MIDI Files. The optional port attribute is a number from 1 to 16 that can be used with the unofficial MIDI port (or cable) meta event. Unlike the DeviceName meta event, there can be multiple midi-device elements per MusicXML part starting in MusicXML 3.0. The optional id attribute refers to the score-instrument assigned to this device. If missing, the device assignment affects all score-instrument elements in the score-part.
type MidiDevice struct {
	XMLName  xml.Name `xml:"midi-device"`
	PortAttr int      `xml:"port,attr,omitempty"`
	IdAttr   string   `xml:"id,attr,omitempty"`
}

// MidiInstrument is The elevation and pan elements allow placing of sound in a 3-D space relative to the listener. Both are expressed in degrees ranging from -180 to 180. For elevation, 0 is level with the listener, 90 is directly above, and -90 is directly below.
type MidiInstrument struct {
	XMLName       xml.Name `xml:"midi-instrument"`
	IdAttr        string   `xml:"id,attr"`
	MidiChannel   int      `xml:"midi-channel"`
	MidiName      string   `xml:"midi-name"`
	MidiBank      int      `xml:"midi-bank"`
	MidiProgram   int      `xml:"midi-program"`
	MidiUnpitched int      `xml:"midi-unpitched"`
	Volume        float64  `xml:"volume"`
	Pan           float64  `xml:"pan"`
	Elevation     float64  `xml:"elevation"`
}

// NameDisplay is The name-display type is used for exact formatting of multi-font text in part and group names to the left of the system. The print-object attribute can be used to determine what, if anything, is printed at the start of each system. Enclosure for the display-text element is none by default. Language for the display-text element is Italian ("it") by default.
type NameDisplay struct {
	XMLName        xml.Name `xml:"name-display"`
	PrintObject    *PrintObject
	DisplayText    *FormattedText  `xml:"display-text"`
	AccidentalText *AccidentalText `xml:"accidental-text"`
}

// OtherPlay is The other-play element represents other types of playback. The required type attribute indicates the type of playback to which the element content applies.
type OtherPlay struct {
	XMLName  xml.Name `xml:"other-play"`
	TypeAttr string   `xml:"type,attr"`
}

// Play is The ipa element represents International Phonetic Alphabet (IPA) sounds for vocal music. String content is limited to IPA 2015 symbols represented in Unicode 13.0.
type Play struct {
	XMLName     xml.Name   `xml:"play"`
	IdAttr      string     `xml:"id,attr,omitempty"`
	Ipa         string     `xml:"ipa"`
	Mute        string     `xml:"mute"`
	SemiPitched string     `xml:"semi-pitched"`
	OtherPlay   *OtherPlay `xml:"other-play"`
}

// Segno is The segno type is the visual indicator of a segno sign. The exact glyph can be specified with the smufl attribute. A sound element is also needed to guide playback applications reliably.
type Segno struct {
	XMLName          xml.Name `xml:"segno"`
	PrintStyleAlign  *PrintStyleAlign
	OptionalUniqueId *OptionalUniqueId
	SmuflAttr        string `xml:"smufl,attr,omitempty"`
}

// String is The string type is used with tablature notation, regular notation (where it is often circled), and chord diagrams. String numbers start with 1 for the highest pitched full-length string.
type String struct {
	XMLName    xml.Name `xml:"string"`
	PrintStyle *PrintStyle
	Placement  *Placement
}

// TypedText is The typed-text type represents a text element with a type attributes.
type TypedText struct {
	XMLName  xml.Name `xml:"typed-text"`
	TypeAttr string   `xml:"type,attr,omitempty"`
}

// WavyLine is Wavy lines are one way to indicate trills. When used with a barline element, they should always have type="continue" set.
type WavyLine struct {
	XMLName    xml.Name `xml:"wavy-line"`
	Position   *Position
	Placement  *Placement
	Color      string
	TrillSound *TrillSound
	TypeAttr   string `xml:"type,attr"`
	NumberAttr int    `xml:"number,attr,omitempty"`
}

// Attributes is The for-part element is used in a concert score to indicate the transposition for a transposed part created from that score. It is only used in score files that contain a concert-score element in the defaults. This allows concert scores with transposed parts to be represented in a single uncompressed MusicXML file.
type Attributes struct {
	XMLName      xml.Name `xml:"attributes"`
	Editorial    *Editorial
	Divisions    float64         `xml:"divisions"`
	Key          []*Key          `xml:"key"`
	Time         []time.Time     `xml:"time"`
	Staves       int             `xml:"staves"`
	PartSymbol   *PartSymbol     `xml:"part-symbol"`
	Instruments  int             `xml:"instruments"`
	Clef         []*Clef         `xml:"clef"`
	StaffDetails []*StaffDetails `xml:"staff-details"`
	Transpose    []*Transpose    `xml:"transpose"`
	ForPart      []*ForPart      `xml:"for-part"`
	Directive    []*Directive    `xml:"directive"`
	MeasureStyle []*MeasureStyle `xml:"measure-style"`
}

// BeatRepeat is The beat-repeat type is used to indicate that a single beat (but possibly many notes) is repeated. The slashes attribute specifies the number of slashes to use in the symbol. The use-dots attribute indicates whether or not to use dots as well (for instance, with mixed rhythm patterns). By default, the value for slashes is 1 and the value for use-dots is no.
//
// The stop attribute indicates the first beat where the repeats are no longer displayed. Both the start and stop of the beat being repeated should be specified unless the repeats are displayed through the end of the part.
//
// The beat-repeat element specifies a notation style for repetitions. The actual music being repeated needs to be repeated within the MusicXML file. This element specifies the notation that indicates the repeat.
type BeatRepeat struct {
	XMLName     xml.Name `xml:"beat-repeat"`
	TypeAttr    string   `xml:"type,attr"`
	SlashesAttr int      `xml:"slashes,attr,omitempty"`
	UseDotsAttr string   `xml:"use-dots,attr,omitempty"`
	Slash       *Slash
}

// Cancel is A cancel element indicates that the old key signature should be cancelled before the new one appears. This will always happen when changing to C major or A minor and need not be specified then. The cancel value matches the fifths value of the cancelled key signature (e.g., a cancel of -2 will provide an explicit cancellation for changing from B flat major to F major). The optional location attribute indicates where the cancellation appears relative to the new key signature.
type Cancel struct {
	XMLName      xml.Name `xml:"cancel"`
	LocationAttr string   `xml:"location,attr,omitempty"`
}

// Clef is Clefs are represented by a combination of sign, line, and clef-octave-change elements. The optional number attribute refers to staff numbers within the part. A value of 1 is assumed if not present.
//
// Sometimes clefs are added to the staff in non-standard line positions, either to indicate cue passages, or when there are multiple clefs present simultaneously on one staff. In this situation, the additional attribute is set to "yes" and the line value is ignored. The size attribute is used for clefs where the additional attribute is "yes". It is typically used to indicate cue clefs.
//
// Sometimes clefs at the start of a measure need to appear after the barline rather than before, as for cues or for use after a repeated section. The after-barline attribute is set to "yes" in this situation. The attribute is ignored for mid-measure clefs.
//
// Clefs appear at the start of each system unless the print-object attribute has been set to "no" or the additional attribute has been set to "yes".
type Clef struct {
	XMLName          xml.Name `xml:"clef"`
	PrintStyle       *PrintStyle
	PrintObject      *PrintObject
	OptionalUniqueId *OptionalUniqueId
	NumberAttr       int    `xml:"number,attr,omitempty"`
	AdditionalAttr   string `xml:"additional,attr,omitempty"`
	SizeAttr         string `xml:"size,attr,omitempty"`
	AfterBarlineAttr string `xml:"after-barline,attr,omitempty"`
	Clef             *Clef
}

// Double is The double type indicates that the music is doubled one octave from what is currently written. If the above attribute is set to yes, the doubling is one octave above what is written, as for mixed flute / piccolo parts in band literature. Otherwise the doubling is one octave below what is written, as for mixed cello / bass parts in orchestral literature.
type Double struct {
	XMLName   xml.Name `xml:"double"`
	AboveAttr string   `xml:"above,attr,omitempty"`
}

// ForPart is The chromatic element in a part-transpose element will usually have a non-zero value, since octave transpositions can be represented in concert scores using the transpose element.
type ForPart struct {
	XMLName       xml.Name       `xml:"for-part"`
	PartClef      *PartClef      `xml:"part-clef"`
	PartTranspose *PartTranspose `xml:"part-transpose"`
}

// Interchangeable is The interchangeable type is used to represent the second in a pair of interchangeable dual time signatures, such as the 6/8 in 3/4 (6/8). A separate symbol attribute value is available compared to the time element's symbol attribute, which applies to the first of the dual time signatures.
type Interchangeable struct {
	XMLName       xml.Name `xml:"interchangeable"`
	SymbolAttr    string   `xml:"symbol,attr,omitempty"`
	SeparatorAttr string   `xml:"separator,attr,omitempty"`
	TimeSignature []*TimeSignature
	TimeRelation  string `xml:"time-relation"`
}

// Key is The optional list of key-octave elements is used to specify in which octave each element of the key signature appears.
type Key struct {
	XMLName           xml.Name `xml:"key"`
	PrintStyle        *PrintStyle
	PrintObject       *PrintObject
	OptionalUniqueId  *OptionalUniqueId
	NumberAttr        int `xml:"number,attr,omitempty"`
	TraditionalKey    *TraditionalKey
	NonTraditionalKey []*NonTraditionalKey
	KeyOctave         []*KeyOctave `xml:"key-octave"`
}

// KeyAccidental is The key-accidental type indicates the accidental to be displayed in a non-traditional key signature, represented in the same manner as the accidental type without the formatting attributes.
type KeyAccidental struct {
	XMLName   xml.Name `xml:"key-accidental"`
	SmuflAttr string   `xml:"smufl,attr,omitempty"`
}

// KeyOctave is The key-octave element specifies in which octave an element of a key signature appears. The content specifies the octave value using the same values as the display-octave element. The number attribute is a positive integer that refers to the key signature element in left-to-right order. If the cancel attribute is set to yes, then this number refers to the canceling key signature specified by the cancel element in the parent key element. The cancel attribute cannot be set to yes if there is no corresponding cancel element within the parent key element. It is no by default.
type KeyOctave struct {
	XMLName    xml.Name `xml:"key-octave"`
	NumberAttr int      `xml:"number,attr"`
	CancelAttr string   `xml:"cancel,attr,omitempty"`
}

// MeasureRepeat is The measure-repeat type is used for both single and multiple measure repeats. The text of the element indicates the number of measures to be repeated in a single pattern. The slashes attribute specifies the number of slashes to use in the repeat sign. It is 1 if not specified. The text of the element is ignored when the type is stop.
//
// The stop attribute indicates the first measure where the repeats are no longer displayed. Both the start and the stop of the measure-repeat should be specified unless the repeats are displayed through the end of the part.
//
// The measure-repeat element specifies a notation style for repetitions. The actual music being repeated needs to be repeated within each measure of the MusicXML file. This element specifies the notation that indicates the repeat.
type MeasureRepeat struct {
	XMLName     xml.Name `xml:"measure-repeat"`
	TypeAttr    string   `xml:"type,attr"`
	SlashesAttr int      `xml:"slashes,attr,omitempty"`
}

// MeasureStyle is A measure-style indicates a special way to print partial to multiple measures within a part. This includes multiple rests over several measures, repeats of beats, single, or multiple measures, and use of slash notation.
//
// The multiple-rest and measure-repeat symbols indicate the number of measures covered in the element content. The beat-repeat and slash elements can cover partial measures. All but the multiple-rest element use a type attribute to indicate starting and stopping the use of the style. The optional number attribute specifies the staff number from top to bottom on the system, as with clef.
type MeasureStyle struct {
	XMLName          xml.Name `xml:"measure-style"`
	Font             *Font
	Color            string
	OptionalUniqueId *OptionalUniqueId
	NumberAttr       int            `xml:"number,attr,omitempty"`
	MultipleRest     *MultipleRest  `xml:"multiple-rest"`
	MeasureRepeat    *MeasureRepeat `xml:"measure-repeat"`
	BeatRepeat       *BeatRepeat    `xml:"beat-repeat"`
	Slash            *Slash         `xml:"slash"`
}

// MultipleRest is The text of the multiple-rest type indicates the number of measures in the multiple rest. Multiple rests may use the 1-bar / 2-bar / 4-bar rest symbols, or a single shape. The use-symbols attribute indicates which to use; it is no if not specified.
type MultipleRest struct {
	XMLName        xml.Name `xml:"multiple-rest"`
	UseSymbolsAttr string   `xml:"use-symbols,attr,omitempty"`
}

// PartClef is The child elements of the part-clef type have the same meaning as for the clef type. However that meaning applies to a transposed part created from the existing score file.
type PartClef struct {
	XMLName xml.Name `xml:"part-clef"`
	Clef    *Clef
}

// PartSymbol is The part-symbol type indicates how a symbol for a multi-staff part is indicated in the score; brace is the default value. The top-staff and bottom-staff elements are used when the brace does not extend across the entire part. For example, in a 3-staff organ part, the top-staff will typically be 1 for the right hand, while the bottom-staff will typically be 2 for the left hand. Staff 3 for the pedals is usually outside the brace.
type PartSymbol struct {
	XMLName         xml.Name `xml:"part-symbol"`
	Position        *Position
	Color           string
	TopStaffAttr    int `xml:"top-staff,attr,omitempty"`
	BottomStaffAttr int `xml:"bottom-staff,attr,omitempty"`
}

// PartTranspose is The child elements of the part-transpose type have the same meaning as for the transpose type. However that meaning applies to a transposed part created from the existing score file.
type PartTranspose struct {
	XMLName   xml.Name `xml:"part-transpose"`
	Transpose *Transpose
}

// Slash is The slash type is used to indicate that slash notation is to be used. If the slash is on every beat, use-stems is no (the default). To indicate rhythms but not pitches, use-stems is set to yes. The type attribute indicates whether this is the start or stop of a slash notation style. The use-dots attribute works as for the beat-repeat element, and only has effect if use-stems is no.
type Slash struct {
	XMLName      xml.Name `xml:"slash"`
	TypeAttr     string   `xml:"type,attr"`
	UseDotsAttr  string   `xml:"use-dots,attr,omitempty"`
	UseStemsAttr string   `xml:"use-stems,attr,omitempty"`
	Slash        *Slash
}

// StaffDetails is The staff-size element indicates how large a staff space is on this staff, expressed as a percentage of the work's default scaling. Values less than 100 make the staff space smaller while values over 100 make the staff space larger. A staff-type of cue, ossia, or editorial implies a staff-size of less than 100, but the exact value is implementation-dependent unless specified here. Staff size affects staff height only, not the relationship of the staff to the left and right margins.
type StaffDetails struct {
	XMLName       xml.Name `xml:"staff-details"`
	PrintObject   *PrintObject
	PrintSpacing  *PrintSpacing
	NumberAttr    int            `xml:"number,attr,omitempty"`
	ShowFretsAttr string         `xml:"show-frets,attr,omitempty"`
	StaffType     string         `xml:"staff-type"`
	StaffLines    int            `xml:"staff-lines"`
	StaffTuning   []*StaffTuning `xml:"staff-tuning"`
	Capo          int            `xml:"capo"`
	StaffSize     float64        `xml:"staff-size"`
}

// StaffTuning is The staff-tuning type specifies the open, non-capo tuning of the lines on a tablature staff.
type StaffTuning struct {
	XMLName  xml.Name `xml:"staff-tuning"`
	LineAttr int      `xml:"line,attr,omitempty"`
	Tuning   *Tuning
}

// Time is A senza-misura element explicitly indicates that no time signature is present. The optional element content indicates the symbol to be used, if any, such as an X. The time element's symbol attribute is not used when a senza-misura element is present.
type Time struct {
	XMLName          xml.Name `xml:"time"`
	PrintStyleAlign  *PrintStyleAlign
	PrintObject      *PrintObject
	OptionalUniqueId *OptionalUniqueId
	NumberAttr       int    `xml:"number,attr,omitempty"`
	SymbolAttr       string `xml:"symbol,attr,omitempty"`
	SeparatorAttr    string `xml:"separator,attr,omitempty"`
	TimeSignature    []*TimeSignature
	Interchangeable  *Interchangeable `xml:"interchangeable"`
	SenzaMisura      string           `xml:"senza-misura"`
}

// Transpose is The transpose type represents what must be added to a written pitch to get a correct sounding pitch. The optional number attribute refers to staff numbers, from top to bottom on the system. If absent, the transposition applies to all staves in the part. Per-staff transposition is most often used in parts that represent multiple instruments.
type Transpose struct {
	XMLName          xml.Name `xml:"transpose"`
	OptionalUniqueId *OptionalUniqueId
	NumberAttr       int `xml:"number,attr,omitempty"`
	Transpose        *Transpose
}

// BarStyleColor is The bar-style-color type contains barline style and color information.
type BarStyleColor struct {
	XMLName xml.Name `xml:"bar-style-color"`
	Color   string
}

// Barline is If a barline is other than a normal single barline, it should be represented by a barline type that describes it. This includes information about repeats and multiple endings, as well as line style. Barline data is on the same level as the other musical data in a score - a child of a measure in a partwise score, or a part in a timewise score. This allows for barlines within measures, as in dotted barlines that subdivide measures in complex meters. The two fermata elements allow for fermatas on both sides of the barline (the lower one inverted).
//
// Barlines have a location attribute to make it easier to process barlines independently of the other musical data in a score. It is often easier to set up measures separately from entering notes. The location attribute must match where the barline element occurs within the rest of the musical data in the score. If location is left, it should be the first element in the measure, aside from the print, bookmark, and link elements. If location is right, it should be the last element, again with the possible exception of the print, bookmark, and link elements. If no location is specified, the right barline is the default. The segno, coda, and divisions attributes work the same way as in the sound element. They are used for playback when barline elements contain segno or coda child elements.
type Barline struct {
	XMLName          xml.Name `xml:"barline"`
	OptionalUniqueId *OptionalUniqueId
	LocationAttr     string  `xml:"location,attr,omitempty"`
	SegnoAttr        string  `xml:"segno,attr,omitempty"`
	CodaAttr         string  `xml:"coda,attr,omitempty"`
	DivisionsAttr    float64 `xml:"divisions,attr,omitempty"`
	Editorial        *Editorial
	BarStyle         *BarStyleColor `xml:"bar-style"`
	WavyLine         *WavyLine      `xml:"wavy-line"`
	Segno            *Segno         `xml:"segno"`
	Coda             *Coda          `xml:"coda"`
	Fermata          []*Fermata     `xml:"fermata"`
	Ending           *Ending        `xml:"ending"`
	Repeat           *Repeat        `xml:"repeat"`
}

// Ending is The ending type represents multiple (e.g. first and second) endings. Typically, the start type is associated with the left barline of the first measure in an ending. The stop and discontinue types are associated with the right barline of the last measure in an ending. Stop is used when the ending mark concludes with a downward jog, as is typical for first endings. Discontinue is used when there is no downward jog, as is typical for second endings that do not conclude a piece. The length of the jog can be specified using the end-length attribute. The text-x and text-y attributes are offsets that specify where the baseline of the start of the ending text appears, relative to the start of the ending line.
//
// The number attribute indicates which times the ending is played, similar to the time-only attribute used by other elements. While this often represents the numeric values for what is under the ending line, it can also indicate whether an ending is played during a larger dal segno or da capo repeat. Single endings such as "1" or comma-separated multiple endings such as "1,2" may be used. The ending element text is used when the text displayed in the ending is different than what appears in the number attribute. The print-object element is used to indicate when an ending is present but not printed, as is often the case for many parts in a full score.
type Ending struct {
	XMLName        xml.Name `xml:"ending"`
	PrintObject    *PrintObject
	PrintStyle     *PrintStyle
	SystemRelation string
	NumberAttr     string  `xml:"number,attr"`
	TypeAttr       string  `xml:"type,attr"`
	EndLengthAttr  float64 `xml:"end-length,attr,omitempty"`
	TextXAttr      float64 `xml:"text-x,attr,omitempty"`
	TextYAttr      float64 `xml:"text-y,attr,omitempty"`
}

// Repeat is The repeat type represents repeat marks. The start of the repeat has a forward direction while the end of the repeat has a backward direction. Backward repeats that are not part of an ending can use the times attribute to indicate the number of times the repeated section is played.
type Repeat struct {
	XMLName       xml.Name `xml:"repeat"`
	DirectionAttr string   `xml:"direction,attr"`
	TimesAttr     int      `xml:"times,attr,omitempty"`
	WingedAttr    string   `xml:"winged,attr,omitempty"`
}

// Accord is The accord type represents the tuning of a single string in the scordatura element. It uses the same group of elements as the staff-tuning element. Strings are numbered from high to low.
type Accord struct {
	XMLName    xml.Name `xml:"accord"`
	StringAttr int      `xml:"string,attr,omitempty"`
	Tuning     *Tuning
}

// AccordionRegistration is The accordion-low element indicates the presence of a dot in the low (16') section of the registration symbol. This element is omitted if no dot is present.
type AccordionRegistration struct {
	XMLName          xml.Name `xml:"accordion-registration"`
	PrintStyleAlign  *PrintStyleAlign
	OptionalUniqueId *OptionalUniqueId
	AccordionHigh    *Empty `xml:"accordion-high"`
	AccordionMiddle  int    `xml:"accordion-middle"`
	AccordionLow     *Empty `xml:"accordion-low"`
}

// Barre is The barre element indicates placing a finger over multiple strings on a single fret. The type is "start" for the lowest pitched string (e.g., the string with the highest MusicXML number) and is "stop" for the highest pitched string.
type Barre struct {
	XMLName  xml.Name `xml:"barre"`
	Color    string
	TypeAttr string `xml:"type,attr"`
}

// Bass is The optional bass-separator element indicates that text, rather than a line or slash, separates the bass from what precedes it.
type Bass struct {
	XMLName         xml.Name   `xml:"bass"`
	ArrangementAttr string     `xml:"arrangement,attr,omitempty"`
	BassSeparator   *StyleText `xml:"bass-separator"`
	BassStep        *BassStep  `xml:"bass-step"`
	BassAlter       *BassAlter `xml:"bass-alter"`
}

// BassAlter is The bass-alter type represents the chromatic alteration of the bass of the current chord within the harmony element. In some chord styles, the text for the bass-step element may include bass-alter information. In that case, the print-object attribute of the bass-alter element can be set to no. The location attribute indicates whether the alteration should appear to the left or the right of the bass-step; it is right by default.
type BassAlter struct {
	XMLName      xml.Name `xml:"bass-alter"`
	PrintObject  *PrintObject
	PrintStyle   *PrintStyle
	LocationAttr string `xml:"location,attr,omitempty"`
}

// BassStep is The bass-step type represents the pitch step of the bass of the current chord within the harmony element. The text attribute indicates how the bass should appear in a score if not using the element contents.
type BassStep struct {
	XMLName    xml.Name `xml:"bass-step"`
	PrintStyle *PrintStyle
	TextAttr   string `xml:"text,attr,omitempty"`
}

// Beater is The beater type represents pictograms for beaters, mallets, and sticks that do not have different materials represented in the pictogram.
type Beater struct {
	XMLName xml.Name `xml:"beater"`
	TipAttr string   `xml:"tip,attr,omitempty"`
}

// BeatUnitTied is The beat-unit-tied type indicates a beat-unit within a metronome mark that is tied to the preceding beat-unit. This allows or two or more tied notes to be associated with a per-minute value in a metronome mark, whereas the metronome-tied element is restricted to metric relationship marks.
type BeatUnitTied struct {
	XMLName  xml.Name `xml:"beat-unit-tied"`
	BeatUnit *BeatUnit
}

// Bracket is Brackets are combined with words in a variety of modern directions. The line-end attribute specifies if there is a jog up or down (or both), an arrow, or nothing at the start or end of the bracket. If the line-end is up or down, the length of the jog can be specified using the end-length attribute. The line-type is solid by default.
type Bracket struct {
	XMLName          xml.Name `xml:"bracket"`
	LineType         string
	DashedFormatting *DashedFormatting
	Position         *Position
	Color            string
	OptionalUniqueId *OptionalUniqueId
	TypeAttr         string  `xml:"type,attr"`
	NumberAttr       int     `xml:"number,attr,omitempty"`
	LineEndAttr      string  `xml:"line-end,attr"`
	EndLengthAttr    float64 `xml:"end-length,attr,omitempty"`
}

// Dashes is The dashes type represents dashes, used for instance with cresc. and dim. marks.
type Dashes struct {
	XMLName          xml.Name `xml:"dashes"`
	DashedFormatting *DashedFormatting
	Position         *Position
	Color            string
	OptionalUniqueId *OptionalUniqueId
	TypeAttr         string `xml:"type,attr"`
	NumberAttr       int    `xml:"number,attr,omitempty"`
}

// Degree is The degree type is used to add, alter, or subtract individual notes in the chord. The print-object attribute can be used to keep the degree from printing separately when it has already taken into account in the text attribute of the kind element. The degree-value and degree-type text attributes specify how the value and type of the degree should be displayed.
//
// A harmony of kind "other" can be spelled explicitly by using a series of degree elements together with a root.
type Degree struct {
	XMLName     xml.Name `xml:"degree"`
	PrintObject *PrintObject
	DegreeValue *DegreeValue `xml:"degree-value"`
	DegreeAlter *DegreeAlter `xml:"degree-alter"`
	DegreeType  *DegreeType  `xml:"degree-type"`
}

// DegreeAlter is The degree-alter type represents the chromatic alteration for the current degree. If the degree-type value is alter or subtract, the degree-alter value is relative to the degree already in the chord based on its kind element. If the degree-type value is add, the degree-alter is relative to a dominant chord (major and perfect intervals except for a minor seventh). The plus-minus attribute is used to indicate if plus and minus symbols should be used instead of sharp and flat symbols to display the degree alteration; it is no by default.
type DegreeAlter struct {
	XMLName       xml.Name `xml:"degree-alter"`
	PrintStyle    *PrintStyle
	PlusMinusAttr string `xml:"plus-minus,attr,omitempty"`
}

// DegreeType is The degree-type type indicates if this degree is an addition, alteration, or subtraction relative to the kind of the current chord. The value of the degree-type element affects the interpretation of the value of the degree-alter element. The text attribute specifies how the type of the degree should be displayed in a score.
type DegreeType struct {
	XMLName    xml.Name `xml:"degree-type"`
	PrintStyle *PrintStyle
	TextAttr   string `xml:"text,attr,omitempty"`
}

// DegreeValue is The content of the degree-value type is a number indicating the degree of the chord (1 for the root, 3 for third, etc). The text attribute specifies how the type of the degree should be displayed in a score. The degree-value symbol attribute indicates that a symbol should be used in specifying the degree. If the symbol attribute is present, the value of the text attribute follows the symbol.
type DegreeValue struct {
	XMLName    xml.Name `xml:"degree-value"`
	PrintStyle *PrintStyle
	SymbolAttr string `xml:"symbol,attr,omitempty"`
	TextAttr   string `xml:"text,attr,omitempty"`
}

// Direction is A direction is a musical indication that is not necessarily attached to a specific note. Two or more may be combined to indicate starts and stops of wedges, dashes, etc. For applications where a specific direction is indeed attached to a specific note, the direction element can be associated with the note element that follows it in score order that is not in a different voice.
//
// By default, a series of direction-type elements and a series of child elements of a direction-type within a single direction element follow one another in sequence visually. For a series of direction-type children, non-positional formatting attributes are carried over from the previous element by default.
type Direction struct {
	XMLName                 xml.Name `xml:"direction"`
	Placement               *Placement
	Directive               *Directive
	SystemRelation          string
	OptionalUniqueId        *OptionalUniqueId
	EditorialVoiceDirection *EditorialVoiceDirection
	Staff                   *Staff
	DirectionType           []*DirectionType `xml:"direction-type"`
	Offset                  *Offset          `xml:"offset"`
	Sound                   *Sound           `xml:"sound"`
	Listening               *Listening       `xml:"listening"`
}

// DirectionType is The eyeglasses element specifies the eyeglasses symbol, common in commercial music.
type DirectionType struct {
	XMLName               xml.Name `xml:"direction-type"`
	OptionalUniqueId      *OptionalUniqueId
	Rehearsal             []*FormattedTextId      `xml:"rehearsal"`
	Segno                 []*Segno                `xml:"segno"`
	Coda                  []*Coda                 `xml:"coda"`
	Words                 *FormattedTextId        `xml:"words"`
	Symbol                *FormattedSymbolId      `xml:"symbol"`
	Wedge                 *Wedge                  `xml:"wedge"`
	Dynamics              []*Dynamics             `xml:"dynamics"`
	Dashes                *Dashes                 `xml:"dashes"`
	Bracket               *Bracket                `xml:"bracket"`
	Pedal                 *Pedal                  `xml:"pedal"`
	Metronome             *Metronome              `xml:"metronome"`
	OctaveShift           *OctaveShift            `xml:"octave-shift"`
	HarpPedals            *HarpPedals             `xml:"harp-pedals"`
	Damp                  *EmptyPrintStyleAlignId `xml:"damp"`
	DampAll               *EmptyPrintStyleAlignId `xml:"damp-all"`
	Eyeglasses            *EmptyPrintStyleAlignId `xml:"eyeglasses"`
	StringMute            *StringMute             `xml:"string-mute"`
	Scordatura            *Scordatura             `xml:"scordatura"`
	Image                 *Image                  `xml:"image"`
	PrincipalVoice        *PrincipalVoice         `xml:"principal-voice"`
	Percussion            []*Percussion           `xml:"percussion"`
	AccordionRegistration *AccordionRegistration  `xml:"accordion-registration"`
	StaffDivide           *StaffDivide            `xml:"staff-divide"`
	OtherDirection        *OtherDirection         `xml:"other-direction"`
}

// Feature is The feature type is a part of the grouping element used for musical analysis. The type attribute represents the type of the feature and the element content represents its value. This type is flexible to allow for different analyses.
type Feature struct {
	XMLName  xml.Name `xml:"feature"`
	TypeAttr string   `xml:"type,attr,omitempty"`
}

// FirstFret is The first-fret type indicates which fret is shown in the top space of the frame; it is fret 1 if the element is not present. The optional text attribute indicates how this is represented in the fret diagram, while the location attribute indicates whether the text appears to the left or right of the frame.
type FirstFret struct {
	XMLName      xml.Name `xml:"first-fret"`
	TextAttr     string   `xml:"text,attr,omitempty"`
	LocationAttr string   `xml:"location,attr,omitempty"`
}

// Frame is The frame-frets element gives the overall size of the frame in horizontal spaces (frets).
type Frame struct {
	XMLName          xml.Name `xml:"frame"`
	Position         *Position
	Color            string
	Halign           *Halign
	ValignImage      string
	OptionalUniqueId *OptionalUniqueId
	HeightAttr       float64      `xml:"height,attr,omitempty"`
	WidthAttr        float64      `xml:"width,attr,omitempty"`
	UnplayedAttr     string       `xml:"unplayed,attr,omitempty"`
	FrameStrings     int          `xml:"frame-strings"`
	FrameFrets       int          `xml:"frame-frets"`
	FirstFret        *FirstFret   `xml:"first-fret"`
	FrameNote        []*FrameNote `xml:"frame-note"`
}

// FrameNote is The frame-note type represents each note included in the frame. An open string will have a fret value of 0, while a muted string will not be associated with a frame-note element.
type FrameNote struct {
	XMLName   xml.Name   `xml:"frame-note"`
	String    string     `xml:"string"`
	Fret      *Fret      `xml:"fret"`
	Fingering *Fingering `xml:"fingering"`
	Barre     *Barre     `xml:"barre"`
}

// Glass is The glass type represents pictograms for glass percussion instruments. The smufl attribute is used to distinguish different SMuFL glyphs for wind chimes in the chimes pictograms range, including those made of materials other than glass.
type Glass struct {
	XMLName   xml.Name `xml:"glass"`
	SmuflAttr string   `xml:"smufl,attr,omitempty"`
}

// Grouping is The grouping type is used for musical analysis. When the type attribute is "start" or "single", it usually contains one or more feature elements. The number attribute is used for distinguishing between overlapping and hierarchical groupings. The member-of attribute allows for easy distinguishing of what grouping elements are in what hierarchy. Feature elements contained within a "stop" type of grouping may be ignored.
//
// This element is flexible to allow for different types of analyses. Future versions of the MusicXML format may add elements that can represent more standardized categories of analysis data, allowing for easier data sharing.
type Grouping struct {
	XMLName          xml.Name `xml:"grouping"`
	OptionalUniqueId *OptionalUniqueId
	TypeAttr         string     `xml:"type,attr"`
	NumberAttr       string     `xml:"number,attr,omitempty"`
	MemberOfAttr     string     `xml:"member-of,attr,omitempty"`
	Feature          []*Feature `xml:"feature"`
}

// Harmony is The harmony type is based on Humdrum's **harm encoding, extended to support chord symbols in popular music as well as functional harmony analysis in classical music.
//
// If there are alternate harmonies possible, this can be specified using multiple harmony elements differentiated by type. Explicit harmonies have all note present in the music; implied have some notes missing but implied; alternate represents alternate analyses.
//
// The harmony object may be used for analysis or for chord symbols. The print-object attribute controls whether or not anything is printed due to the harmony element. The print-frame attribute controls printing of a frame or fretboard diagram. The print-style attribute group sets the default for the harmony, but individual elements can override this with their own print-style values. The arrangement attribute specifies how multiple harmony-chord groups are arranged relative to each other.
type Harmony struct {
	XMLName          xml.Name `xml:"harmony"`
	PrintObject      *PrintObject
	PrintStyle       *PrintStyle
	Placement        *Placement
	SystemRelation   string
	OptionalUniqueId *OptionalUniqueId
	TypeAttr         string `xml:"type,attr,omitempty"`
	PrintFrameAttr   string `xml:"print-frame,attr,omitempty"`
	ArrangementAttr  string `xml:"arrangement,attr,omitempty"`
	HarmonyChord     []*HarmonyChord
	Editorial        *Editorial
	Staff            *Staff
	Frame            *Frame  `xml:"frame"`
	Offset           *Offset `xml:"offset"`
}

// HarpPedals is The harp-pedals type is used to create harp pedal diagrams. The pedal-step and pedal-alter elements use the same values as the step and alter elements. For easiest reading, the pedal-tuning elements should follow standard harp pedal order, with pedal-step values of D, C, B, E, F, G, and A.
type HarpPedals struct {
	XMLName          xml.Name `xml:"harp-pedals"`
	PrintStyleAlign  *PrintStyleAlign
	OptionalUniqueId *OptionalUniqueId
	PedalTuning      []*PedalTuning `xml:"pedal-tuning"`
}

// Image is The image type is used to include graphical images in a score.
type Image struct {
	XMLName          xml.Name `xml:"image"`
	ImageAttributes  *ImageAttributes
	OptionalUniqueId *OptionalUniqueId
}

// InstrumentChange is The instrument-change element type represents a change to the virtual instrument sound for a given score-instrument. The id attribute refers to the score-instrument affected by the change. All members of the instrument-change type can also be initially specified within the score-instrument element.
type InstrumentChange struct {
	XMLName               xml.Name `xml:"instrument-change"`
	IdAttr                string   `xml:"id,attr"`
	VirtualInstrumentData *VirtualInstrumentData
}

// Inversion is The inversion type represents harmony inversions. The value is a number indicating which inversion is used: 0 for root position, 1 for first inversion, etc.
type Inversion struct {
	XMLName    xml.Name `xml:"inversion"`
	PrintStyle *PrintStyle
}

// Kind is Kind indicates the type of chord. Degree elements can then add, subtract, or alter from these starting points
//
// The attributes are used to indicate the formatting of the symbol. Since the kind element is the constant in all the harmony-chord groups that can make up a polychord, many formatting attributes are here.
//
// The use-symbols attribute is yes if the kind should be represented when possible with harmony symbols rather than letters and numbers. These symbols include:
//
// major: a triangle, like Unicode 25B3
// minor: -, like Unicode 002D
// augmented: +, like Unicode 002B
// diminished: °, like Unicode 00B0
// half-diminished: ø, like Unicode 00F8
//
// For the major-minor kind, only the minor symbol is used when use-symbols is yes. The major symbol is set using the symbol attribute in the degree-value element. The corresponding degree-alter value will usually be 0 in this case.
//
// The text attribute describes how the kind should be spelled in a score. If use-symbols is yes, the value of the text attribute follows the symbol. The stack-degrees attribute is yes if the degree elements should be stacked above each other. The parentheses-degrees attribute is yes if all the degrees should be in parentheses. The bracket-degrees attribute is yes if all the degrees should be in a bracket. If not specified, these values are implementation-specific. The alignment attributes are for the entire harmony-chord group of which this kind element is a part.
//
// The text attribute may use strings such as "13sus" that refer to both the kind and one or more degree elements. In this case, the corresponding degree elements should have the print-object attribute set to "no" to keep redundant alterations from being displayed.
type Kind struct {
	XMLName                xml.Name `xml:"kind"`
	PrintStyle             *PrintStyle
	Halign                 *Halign
	Valign                 string
	UseSymbolsAttr         string `xml:"use-symbols,attr,omitempty"`
	TextAttr               string `xml:"text,attr,omitempty"`
	StackDegreesAttr       string `xml:"stack-degrees,attr,omitempty"`
	ParenthesesDegreesAttr string `xml:"parentheses-degrees,attr,omitempty"`
	BracketDegreesAttr     string `xml:"bracket-degrees,attr,omitempty"`
}

// Listening is The listen and listening types, new in Version 4.0, specify different ways that a score following or machine listening application can interact with a performer. The listening type handles interactions that change the state of the listening application from the specified point in the performance onward. If multiple child elements of the same type are present, they should have distinct player and/or time-only attributes.
//
// The offset element is used to indicate that the listening change takes place offset from the current score position. If the listening element is a child of a direction element, the listening offset element overrides the direction offset element if both elements are present. Note that the offset reflects the intended musical position for the change in state. It should not be used to compensate for latency issues in particular hardware configurations.
type Listening struct {
	XMLName        xml.Name        `xml:"listening"`
	Sync           *Sync           `xml:"sync"`
	OtherListening *OtherListening `xml:"other-listening"`
	Offset         *Offset         `xml:"offset"`
}

// MeasureNumbering is The measure-numbering type describes how frequently measure numbers are displayed on this part. The text attribute from the measure element is used for display, or the number attribute if the text attribute is not present. Measures with an implicit attribute set to "yes" never display a measure number, regardless of the measure-numbering setting.
//
// The optional staff attribute refers to staff numbers within the part, from top to bottom on the system. It indicates which staff is used as the reference point for vertical positioning. A value of 1 is assumed if not present.
//
// The optional multiple-rest-always and multiple-rest-range attributes describe how measure numbers are shown on multiple rests when the measure-numbering value is not set to none. The multiple-rest-always attribute is set to yes when the measure number should always be shown, even if the multiple rest starts midway through a system when measure numbering is set to system level. The multiple-rest-range attribute is set to yes when measure numbers on multiple rests display the range of numbers for the first and last measure, rather than just the number of the first measure.
type MeasureNumbering struct {
	XMLName                xml.Name `xml:"measure-numbering"`
	PrintStyleAlign        *PrintStyleAlign
	SystemAttr             string `xml:"system,attr,omitempty"`
	StaffAttr              int    `xml:"staff,attr,omitempty"`
	MultipleRestAlwaysAttr string `xml:"multiple-rest-always,attr,omitempty"`
	MultipleRestRangeAttr  string `xml:"multiple-rest-range,attr,omitempty"`
}

// Metronome is The metronome-relation element describes the relationship symbol that goes between the two sets of metronome-note elements. The currently allowed value is equals, but this may expand in future versions. If the element is empty, the equals value is used.
type Metronome struct {
	XMLName           xml.Name `xml:"metronome"`
	PrintStyleAlign   *PrintStyleAlign
	Justify           *Justify
	OptionalUniqueId  *OptionalUniqueId
	ParenthesesAttr   string `xml:"parentheses,attr,omitempty"`
	BeatUnit          *BeatUnit
	BeatUnitTied      []*BeatUnitTied  `xml:"beat-unit-tied"`
	PerMinute         *PerMinute       `xml:"per-minute"`
	MetronomeArrows   *Empty           `xml:"metronome-arrows"`
	MetronomeNote     []*MetronomeNote `xml:"metronome-note"`
	MetronomeRelation string           `xml:"metronome-relation"`
}

// MetronomeBeam is The metronome-beam type works like the beam type in defining metric relationships, but does not include all the attributes available in the beam type.
type MetronomeBeam struct {
	XMLName    xml.Name `xml:"metronome-beam"`
	NumberAttr int      `xml:"number,attr,omitempty"`
}

// MetronomeNote is The metronome-dot element works like the dot element in defining metric relationships.
type MetronomeNote struct {
	XMLName         xml.Name         `xml:"metronome-note"`
	MetronomeType   string           `xml:"metronome-type"`
	MetronomeDot    []*Empty         `xml:"metronome-dot"`
	MetronomeBeam   []*MetronomeBeam `xml:"metronome-beam"`
	MetronomeTied   *MetronomeTied   `xml:"metronome-tied"`
	MetronomeTuplet *MetronomeTuplet `xml:"metronome-tuplet"`
}

// MetronomeTied is The metronome-tied indicates the presence of a tie within a metric relationship mark. As with the tied element, both the start and stop of the tie should be specified, in this case within separate metronome-note elements.
type MetronomeTied struct {
	XMLName  xml.Name `xml:"metronome-tied"`
	TypeAttr string   `xml:"type,attr"`
}

// MetronomeTuplet is The metronome-tuplet type uses the same element structure as the time-modification element along with some attributes from the tuplet element.
type MetronomeTuplet struct {
	XMLName        xml.Name `xml:"metronome-tuplet"`
	TypeAttr       string   `xml:"type,attr"`
	BracketAttr    string   `xml:"bracket,attr,omitempty"`
	ShowNumberAttr string   `xml:"show-number,attr,omitempty"`
}

// OctaveShift is The octave shift type indicates where notes are shifted up or down from their true pitched values because of printing difficulty. Thus a treble clef line noted with 8va will be indicated with an octave-shift down from the pitch data indicated in the notes. A size of 8 indicates one octave; a size of 15 indicates two octaves.
type OctaveShift struct {
	XMLName          xml.Name `xml:"octave-shift"`
	DashedFormatting *DashedFormatting
	PrintStyle       *PrintStyle
	OptionalUniqueId *OptionalUniqueId
	TypeAttr         string `xml:"type,attr"`
	NumberAttr       int    `xml:"number,attr,omitempty"`
	SizeAttr         int    `xml:"size,attr,omitempty"`
}

// Offset is An offset is represented in terms of divisions, and indicates where the direction will appear relative to the current musical location. The current musical location is always within the current measure, even at the end of a measure.
//
// The offset affects the visual appearance of the direction. If the sound attribute is "yes", then the offset affects playback and listening too. If the sound attribute is "no", then any sound or listening associated with the direction takes effect at the current location. The sound attribute is "no" by default for compatibility with earlier versions of the MusicXML format. If an element within a direction includes a default-x attribute, the offset value will be ignored when determining the appearance of that element.
type Offset struct {
	XMLName   xml.Name `xml:"offset"`
	SoundAttr string   `xml:"sound,attr,omitempty"`
}

// OtherDirection is The other-direction type is used to define any direction symbols not yet in the MusicXML format. The smufl attribute can be used to specify a particular direction symbol, allowing application interoperability without requiring every SMuFL glyph to have a MusicXML element equivalent. Using the other-direction type without the smufl attribute allows for extended representation, though without application interoperability.
type OtherDirection struct {
	XMLName          xml.Name `xml:"other-direction"`
	PrintObject      *PrintObject
	PrintStyleAlign  *PrintStyleAlign
	Smufl            *Smufl
	OptionalUniqueId *OptionalUniqueId
}

// OtherListening is TThe other-listening type represents other types of listening control and interaction. The required type attribute indicates the type of listening to which the element content applies. The optional player and time-only attributes restrict the element to apply to a single player or set of times through a repeated section, respectively.
type OtherListening struct {
	XMLName      xml.Name `xml:"other-listening"`
	TypeAttr     string   `xml:"type,attr"`
	PlayerAttr   string   `xml:"player,attr,omitempty"`
	TimeOnlyAttr string   `xml:"time-only,attr,omitempty"`
}

// Pedal is The pedal type represents piano pedal marks. Starting with MusicXML 3.1 this includes sostenuto as well as damper pedal marks. The line attribute is yes if pedal lines are used. The sign attribute is yes if Ped, Sost, and * signs are used. For MusicXML 2.0 compatibility, the sign attribute is yes by default if the line attribute is no, and is no by default if the line attribute is yes. If the sign attribute is set to yes and the type is start or sostenuto, the abbreviated attribute is yes if the short P and S signs are used, and no if the full Ped and Sost signs are used. It is no by default. Otherwise the abbreviated attribute is ignored. The alignment attributes are ignored if the line attribute is yes.
type Pedal struct {
	XMLName          xml.Name `xml:"pedal"`
	PrintStyleAlign  *PrintStyleAlign
	OptionalUniqueId *OptionalUniqueId
	TypeAttr         string `xml:"type,attr"`
	NumberAttr       int    `xml:"number,attr,omitempty"`
	LineAttr         string `xml:"line,attr,omitempty"`
	SignAttr         string `xml:"sign,attr,omitempty"`
	AbbreviatedAttr  string `xml:"abbreviated,attr,omitempty"`
}

// PedalTuning is The pedal-alter element defines the chromatic alteration for a single harp pedal.
type PedalTuning struct {
	XMLName    xml.Name `xml:"pedal-tuning"`
	PedalStep  string   `xml:"pedal-step"`
	PedalAlter float64  `xml:"pedal-alter"`
}

// PerMinute is The per-minute type can be a number, or a text description including numbers. If a font is specified, it overrides the font specified for the overall metronome element. This allows separate specification of a music font for the beat-unit and a text font for the numeric value, in cases where a single metronome font is not used.
type PerMinute struct {
	XMLName xml.Name `xml:"per-minute"`
	Font    *Font
}

// Percussion is The percussion element is used to define percussion pictogram symbols. Definitions for these symbols can be found in Kurt Stone's "Music Notation in the Twentieth Century" on pages 206-212 and 223. Some values are added to these based on how usage has evolved in the 30 years since Stone's book was published.
type Percussion struct {
	XMLName          xml.Name `xml:"percussion"`
	PrintStyleAlign  *PrintStyleAlign
	Enclosure        *Enclosure
	OptionalUniqueId *OptionalUniqueId
	Glass            *Glass     `xml:"glass"`
	Metal            string     `xml:"metal"`
	Wood             string     `xml:"wood"`
	Pitched          *Pitched   `xml:"pitched"`
	Membrane         string     `xml:"membrane"`
	Effect           string     `xml:"effect"`
	Timpani          *Empty     `xml:"timpani"`
	Beater           *Beater    `xml:"beater"`
	Stick            *Stick     `xml:"stick"`
	StickLocation    string     `xml:"stick-location"`
	OtherPercussion  *OtherText `xml:"other-percussion"`
}

// Pitched is The pitched-value type represents pictograms for pitched percussion instruments. The smufl attribute is used to distinguish different SMuFL glyphs for a particular pictogram within the tuned mallet percussion pictograms range.
type Pitched struct {
	XMLName   xml.Name `xml:"pitched"`
	SmuflAttr string   `xml:"smufl,attr,omitempty"`
}

// PrincipalVoice is The principal-voice element represents principal and secondary voices in a score, either for analysis or for square bracket symbols that appear in a score. The symbol attribute indicates the type of symbol used at the start of the principal-voice. The content of the principal-voice element is used for analysis and may be any text value. When used for analysis separate from any printed score markings, the symbol attribute should be set to "none".
type PrincipalVoice struct {
	XMLName          xml.Name `xml:"principal-voice"`
	PrintStyleAlign  *PrintStyleAlign
	OptionalUniqueId *OptionalUniqueId
	TypeAttr         string `xml:"type,attr"`
	SymbolAttr       string `xml:"symbol,attr"`
}

// Print is The print type contains general printing parameters, including the layout elements defined in the layout.mod file. The part-name-display and part-abbreviation-display elements used in the score.mod file may also be used here to change how a part name or abbreviation is displayed over the course of a piece. They take effect when the current measure or a succeeding measure starts a new system.
//
// Layout group elements in a print statement only apply to the current page, system, or staff. Music that follows continues to take the default values from the layout determined by the defaults element.
type Print struct {
	XMLName                 xml.Name `xml:"print"`
	PrintAttributes         *PrintAttributes
	OptionalUniqueId        *OptionalUniqueId
	Layout                  *Layout
	MeasureLayout           *MeasureLayout    `xml:"measure-layout"`
	MeasureNumbering        *MeasureNumbering `xml:"measure-numbering"`
	PartNameDisplay         *NameDisplay      `xml:"part-name-display"`
	PartAbbreviationDisplay *NameDisplay      `xml:"part-abbreviation-display"`
}

// Root is The root type indicates a pitch like C, D, E vs. a function indication like I, II, III. It is used with chord symbols in popular music. The root element has a root-step and optional root-alter element similar to the step and alter elements, but renamed to distinguish the different musical meanings.
type Root struct {
	XMLName   xml.Name   `xml:"root"`
	RootStep  *RootStep  `xml:"root-step"`
	RootAlter *RootAlter `xml:"root-alter"`
}

// RootAlter is The root-alter type represents the chromatic alteration of the root of the current chord within the harmony element. In some chord styles, the text for the root-step element may include root-alter information. In that case, the print-object attribute of the root-alter element can be set to no. The location attribute indicates whether the alteration should appear to the left or the right of the root-step; it is right by default.
type RootAlter struct {
	XMLName      xml.Name `xml:"root-alter"`
	PrintObject  *PrintObject
	PrintStyle   *PrintStyle
	LocationAttr string `xml:"location,attr,omitempty"`
}

// RootStep is The root-step type represents the pitch step of the root of the current chord within the harmony element. The text attribute indicates how the root should appear in a score if not using the element contents.
type RootStep struct {
	XMLName    xml.Name `xml:"root-step"`
	PrintStyle *PrintStyle
	TextAttr   string `xml:"text,attr,omitempty"`
}

// Scordatura is Scordatura string tunings are represented by a series of accord elements, similar to the staff-tuning elements. Strings are numbered from high to low.
type Scordatura struct {
	XMLName          xml.Name `xml:"scordatura"`
	OptionalUniqueId *OptionalUniqueId
	Accord           []*Accord `xml:"accord"`
}

// Sound is The sound element contains general playback parameters. They can stand alone within a part/measure, or be a component element within a direction.
//
// Tempo is expressed in quarter notes per minute. If 0, the sound-generating program should prompt the user at the time of compiling a sound (MIDI) file.
//
// Dynamics (or MIDI velocity) are expressed as a percentage of the default forte value (90 for MIDI 1.0).
//
// Dacapo indicates to go back to the beginning of the movement. When used it always has the value "yes".
//
// Segno and dalsegno are used for backwards jumps to a segno sign; coda and tocoda are used for forward jumps to a coda sign. If there are multiple jumps, the value of these parameters can be used to name and distinguish them. If segno or coda is used, the divisions attribute can also be used to indicate the number of divisions per quarter note. Otherwise sound and MIDI generating programs may have to recompute this.
//
// By default, a dalsegno or dacapo attribute indicates that the jump should occur the first time through, while a tocoda attribute indicates the jump should occur the second time through. The time that jumps occur can be changed by using the time-only attribute.
//
// Forward-repeat is used when a forward repeat sign is implied, and usually follows a bar line. When used it always has the value of "yes".
//
// The fine attribute follows the final note or rest in a movement with a da capo or dal segno direction. If numeric, the value represents the actual duration of the final note or rest, which can be ambiguous in written notation and different among parts and voices. The value may also be "yes" to indicate no change to the final duration.
//
// If the sound element applies only particular times through a repeat, the time-only attribute indicates which times to apply the sound element.
//
// Pizzicato in a sound element effects all following notes. Yes indicates pizzicato, no indicates arco.
//
// The pan and elevation attributes are deprecated in Version 2.0. The pan and elevation elements in the midi-instrument element should be used instead. The meaning of the pan and elevation attributes is the same as for the pan and elevation elements. If both are present, the mid-instrument elements take priority.
//
// The damper-pedal, soft-pedal, and sostenuto-pedal attributes effect playback of the three common piano pedals and their MIDI controller equivalents. The yes value indicates the pedal is depressed; no indicates the pedal is released. A numeric value from 0 to 100 may also be used for half pedaling. This value is the percentage that the pedal is depressed. A value of 0 is equivalent to no, and a value of 100 is equivalent to yes.
//
// Instrument changes, MIDI devices, MIDI instruments, and playback techniques are changed using the instrument-change, midi-device, midi-instrument, and play elements. When there are multiple instances of these elements, they should be grouped together by instrument using the id attribute values.
//
// The offset element is used to indicate that the sound takes place offset from the current score position. If the sound element is a child of a direction element, the sound offset element overrides the direction offset element if both elements are present. Note that the offset reflects the intended musical position for the change in sound. It should not be used to compensate for latency issues in particular hardware configurations.
type Sound struct {
	XMLName            xml.Name `xml:"sound"`
	OptionalUniqueId   *OptionalUniqueId
	TempoAttr          float64           `xml:"tempo,attr,omitempty"`
	DynamicsAttr       float64           `xml:"dynamics,attr,omitempty"`
	DacapoAttr         string            `xml:"dacapo,attr,omitempty"`
	SegnoAttr          string            `xml:"segno,attr,omitempty"`
	DalsegnoAttr       string            `xml:"dalsegno,attr,omitempty"`
	CodaAttr           string            `xml:"coda,attr,omitempty"`
	TocodaAttr         string            `xml:"tocoda,attr,omitempty"`
	DivisionsAttr      float64           `xml:"divisions,attr,omitempty"`
	ForwardRepeatAttr  string            `xml:"forward-repeat,attr,omitempty"`
	FineAttr           string            `xml:"fine,attr,omitempty"`
	TimeOnlyAttr       string            `xml:"time-only,attr,omitempty"`
	PizzicatoAttr      string            `xml:"pizzicato,attr,omitempty"`
	PanAttr            float64           `xml:"pan,attr,omitempty"`
	ElevationAttr      float64           `xml:"elevation,attr,omitempty"`
	DamperPedalAttr    *YesNoNumber      `xml:"damper-pedal,attr,omitempty"`
	SoftPedalAttr      *YesNoNumber      `xml:"soft-pedal,attr,omitempty"`
	SostenutoPedalAttr *YesNoNumber      `xml:"sostenuto-pedal,attr,omitempty"`
	InstrumentChange   *InstrumentChange `xml:"instrument-change"`
	MidiDevice         *MidiDevice       `xml:"midi-device"`
	MidiInstrument     *MidiInstrument   `xml:"midi-instrument"`
	Play               *Play             `xml:"play"`
	Swing              *Swing            `xml:"swing"`
	Offset             *Offset           `xml:"offset"`
}

// StaffDivide is The staff-divide element represents the staff division arrow symbols found at SMuFL code points U+E00B, U+E00C, and U+E00D.
type StaffDivide struct {
	XMLName          xml.Name `xml:"staff-divide"`
	PrintStyleAlign  *PrintStyleAlign
	OptionalUniqueId *OptionalUniqueId
	TypeAttr         string `xml:"type,attr"`
}

// Stick is The stick type represents pictograms where the material of the stick, mallet, or beater is included.The parentheses and dashed-circle attributes indicate the presence of these marks around the round beater part of a pictogram. Values for these attributes are "no" if not present.
type Stick struct {
	XMLName          xml.Name `xml:"stick"`
	TipAttr          string   `xml:"tip,attr,omitempty"`
	ParenthesesAttr  string   `xml:"parentheses,attr,omitempty"`
	DashedCircleAttr string   `xml:"dashed-circle,attr,omitempty"`
	StickType        string   `xml:"stick-type"`
	StickMaterial    string   `xml:"stick-material"`
}

// StringMute is The string-mute type represents string mute on and mute off symbols.
type StringMute struct {
	XMLName          xml.Name `xml:"string-mute"`
	PrintStyleAlign  *PrintStyleAlign
	OptionalUniqueId *OptionalUniqueId
	TypeAttr         string `xml:"type,attr"`
}

// Swing is The swing element specifies whether or not to use swing playback, where consecutive on-beat / off-beat eighth or 16th notes are played with unequal nominal durations.
//
// The straight element specifies that no swing is present, so consecutive notes have equal durations.
//
// The first and second elements are positive integers that specify the ratio between durations of consecutive notes. For example, a first element with a value of 2 and a second element with a value of 1 applied to eighth notes specifies a quarter note / eighth note tuplet playback, where the first note is twice as long as the second note. Ratios should be specified with the smallest integers possible. For example, a ratio of 6 to 4 should be specified as 3 to 2 instead.
//
// The optional swing-type element specifies the note type, either eighth or 16th, to which the ratio is applied. The value is eighth if this element is not present.
//
// The optional swing-style element is a string describing the style of swing used.
//
// The swing element has no effect for playback of grace notes, notes where a type element is not present, and notes where the specified duration is different than the nominal value associated with the specified type. If a swung note has attack and release attributes, those values modify the swung playback.
type Swing struct {
	XMLName    xml.Name `xml:"swing"`
	Straight   *Empty   `xml:"straight"`
	First      int      `xml:"first"`
	Second     int      `xml:"second"`
	SwingType  string   `xml:"swing-type"`
	SwingStyle string   `xml:"swing-style"`
}

// Sync is The sync type specifies the style that a score following application should use the synchronize an accompaniment with a performer. If this type is not included in a score, default synchronization depends on the application.
//
// The optional latency attribute specifies a time in milliseconds that the listening application should expect from the performer. The optional player and time-only attributes restrict the element to apply to a single player or set of times through a repeated section, respectively.
type Sync struct {
	XMLName      xml.Name `xml:"sync"`
	TypeAttr     string   `xml:"type,attr"`
	LatencyAttr  int      `xml:"latency,attr,omitempty"`
	PlayerAttr   string   `xml:"player,attr,omitempty"`
	TimeOnlyAttr string   `xml:"time-only,attr,omitempty"`
}

// Wedge is The wedge type represents crescendo and diminuendo wedge symbols. The type attribute is crescendo for the start of a wedge that is closed at the left side, and diminuendo for the start of a wedge that is closed on the right side. Spread values are measured in tenths; those at the start of a crescendo wedge or end of a diminuendo wedge are ignored. The niente attribute is yes if a circle appears at the point of the wedge, indicating a crescendo from nothing or diminuendo to nothing. It is no by default, and used only when the type is crescendo, or the type is stop for a wedge that began with a diminuendo type. The line-type is solid by default.
type Wedge struct {
	XMLName          xml.Name `xml:"wedge"`
	LineType         string
	DashedFormatting *DashedFormatting
	Position         *Position
	Color            string
	OptionalUniqueId *OptionalUniqueId
	TypeAttr         string  `xml:"type,attr"`
	NumberAttr       int     `xml:"number,attr,omitempty"`
	SpreadAttr       float64 `xml:"spread,attr,omitempty"`
	NienteAttr       string  `xml:"niente,attr,omitempty"`
}

// Encoding is The encoding element contains information about who did the digital encoding, when, with what software, and in what aspects. Standard type values for the encoder element are music, words, and arrangement, but other types may be used. The type attribute is only needed when there are multiple encoder elements.
type Encoding struct {
	XMLName             xml.Name   `xml:"encoding"`
	EncodingDate        time.Time  `xml:"encoding-date"`
	Encoder             *TypedText `xml:"encoder"`
	Software            string     `xml:"software"`
	EncodingDescription string     `xml:"encoding-description"`
	Supports            *Supports  `xml:"supports"`
}

// Identification is A related resource for the music that is encoded. This is similar to the Dublin Core relation element. Standard type values are music, words, and arrangement, but other types may be used.
type Identification struct {
	XMLName       xml.Name       `xml:"identification"`
	Creator       []*TypedText   `xml:"creator"`
	Rights        []*TypedText   `xml:"rights"`
	Encoding      *Encoding      `xml:"encoding"`
	Source        string         `xml:"source"`
	Relation      []*TypedText   `xml:"relation"`
	Miscellaneous *Miscellaneous `xml:"miscellaneous"`
}

// Miscellaneous is If a program has other metadata not yet supported in the MusicXML format, it can go in the miscellaneous element. The miscellaneous type puts each separate part of metadata into its own miscellaneous-field type.
type Miscellaneous struct {
	XMLName            xml.Name              `xml:"miscellaneous"`
	MiscellaneousField []*MiscellaneousField `xml:"miscellaneous-field"`
}

// MiscellaneousField is If a program has other metadata not yet supported in the MusicXML format, each type of metadata can go in a miscellaneous-field element. The required name attribute indicates the type of metadata the element content represents.
type MiscellaneousField struct {
	XMLName  xml.Name `xml:"miscellaneous-field"`
	NameAttr string   `xml:"name,attr"`
}

// Supports is The supports type indicates if a MusicXML encoding supports a particular MusicXML element. This is recommended for elements like beam, stem, and accidental, where the absence of an element is ambiguous if you do not know if the encoding supports that element. For Version 2.0, the supports element is expanded to allow programs to indicate support for particular attributes or particular values. This lets applications communicate, for example, that all system and/or page breaks are contained in the MusicXML file.
type Supports struct {
	XMLName       xml.Name `xml:"supports"`
	TypeAttr      string   `xml:"type,attr"`
	ElementAttr   string   `xml:"element,attr"`
	AttributeAttr string   `xml:"attribute,attr,omitempty"`
	ValueAttr     string   `xml:"value,attr,omitempty"`
}

// Appearance is The appearance type controls general graphical settings for the music's final form appearance on a printed page of display. This includes support for line widths, definitions for note sizes, and standard distances between notation elements, plus an extension element for other aspects of appearance.
type Appearance struct {
	XMLName         xml.Name           `xml:"appearance"`
	LineWidth       []*LineWidth       `xml:"line-width"`
	NoteSize        []*NoteSize        `xml:"note-size"`
	Distance        []*Distance        `xml:"distance"`
	Glyph           []*Glyph           `xml:"glyph"`
	OtherAppearance []*OtherAppearance `xml:"other-appearance"`
}

// Distance is The distance element represents standard distances between notation elements in tenths. The type attribute defines what type of distance is being defined. Valid values include hyphen (for hyphens in lyrics) and beam.
type Distance struct {
	XMLName  xml.Name `xml:"distance"`
	TypeAttr string   `xml:"type,attr"`
}

// Glyph is The glyph element represents what SMuFL glyph should be used for different variations of symbols that are semantically identical. The type attribute specifies what type of glyph is being defined. The element value specifies what SMuFL glyph to use, including recommended stylistic alternates. The SMuFL glyph name should match the type. For instance, a type of quarter-rest would use values restQuarter, restQuarterOld, or restQuarterZ. A type of g-clef-ottava-bassa would use values gClef8vb, gClef8vbOld, or gClef8vbCClef. A type of octave-shift-up-8 would use values ottava, ottavaBassa, ottavaBassaBa, ottavaBassaVb, or octaveBassa.
type Glyph struct {
	XMLName  xml.Name `xml:"glyph"`
	TypeAttr string   `xml:"type,attr"`
}

// LineWidth is The line-width type indicates the width of a line type in tenths. The type attribute defines what type of line is being defined. Values include beam, bracket, dashes, enclosure, ending, extend, heavy barline, leger, light barline, octave shift, pedal, slur middle, slur tip, staff, stem, tie middle, tie tip, tuplet bracket, and wedge. The text content is expressed in tenths.
type LineWidth struct {
	XMLName  xml.Name `xml:"line-width"`
	TypeAttr string   `xml:"type,attr"`
}

// MeasureLayout is The measure-distance element specifies the horizontal distance from the previous measure. This value is only used for systems where there is horizontal whitespace in the middle of a system, as in systems with codas. To specify the measure width, use the width attribute of the measure element.
type MeasureLayout struct {
	XMLName         xml.Name `xml:"measure-layout"`
	MeasureDistance float64  `xml:"measure-distance"`
}

// NoteSize is The note-size type indicates the percentage of the regular note size to use for notes with a cue and large size as defined in the type element. The grace type is used for notes of cue size that that include a grace element. The cue type is used for all other notes with cue size, whether defined explicitly or implicitly via a cue element. The large type is used for notes of large size. The text content represent the numeric percentage. A value of 100 would be identical to the size of a regular note as defined by the music font.
type NoteSize struct {
	XMLName  xml.Name `xml:"note-size"`
	TypeAttr string   `xml:"type,attr"`
}

// OtherAppearance is The other-appearance type is used to define any graphical settings not yet in the current version of the MusicXML format. This allows extended representation, though without application interoperability.
type OtherAppearance struct {
	XMLName  xml.Name `xml:"other-appearance"`
	TypeAttr string   `xml:"type,attr"`
}

// PageLayout is Page layout can be defined both in score-wide defaults and in the print element. Page margins are specified either for both even and odd pages, or via separate odd and even page number values. The type is not needed when used as part of a print element. If omitted when used in the defaults element, "both" is the default.
//
// If no page-layout element is present in the defaults element, default page layout values are chosen by the application.
//
// When used in the print element, the page-layout element affects the appearance of the current page only. All other pages use the default values as determined by the defaults element. If any child elements are missing from the page-layout element in a print element, the values determined by the defaults element are used there as well.
type PageLayout struct {
	XMLName     xml.Name       `xml:"page-layout"`
	PageHeight  float64        `xml:"page-height"`
	PageWidth   float64        `xml:"page-width"`
	PageMargins []*PageMargins `xml:"page-margins"`
}

// PageMargins is Page margins are specified either for both even and odd pages, or via separate odd and even page number values. The type attribute is not needed when used as part of a print element. If omitted when the page-margins type is used in the defaults element, "both" is the default value.
type PageMargins struct {
	XMLName    xml.Name `xml:"page-margins"`
	TypeAttr   string   `xml:"type,attr,omitempty"`
	AllMargins *AllMargins
}

// Scaling is Margins, page sizes, and distances are all measured in tenths to keep MusicXML data in a consistent coordinate system as much as possible. The translation to absolute units is done with the scaling type, which specifies how many millimeters are equal to how many tenths. For a staff height of 7 mm, millimeters would be set to 7 while tenths is set to 40. The ability to set a formula rather than a single scaling factor helps avoid roundoff errors.
type Scaling struct {
	XMLName     xml.Name `xml:"scaling"`
	Millimeters float64  `xml:"millimeters"`
	Tenths      float64  `xml:"tenths"`
}

// StaffLayout is Staff layout includes the vertical distance from the bottom line of the previous staff in this system to the top line of the staff specified by the number attribute. The optional number attribute refers to staff numbers within the part, from top to bottom on the system. A value of 1 is assumed if not present.
//
// When used in the defaults element, the values apply to all systems in all parts. When used in the print element, the values apply to the current system only. This value is ignored for the first staff in a system.
type StaffLayout struct {
	XMLName       xml.Name `xml:"staff-layout"`
	NumberAttr    int      `xml:"number,attr,omitempty"`
	StaffDistance float64  `xml:"staff-distance"`
}

// SystemDividers is The system-dividers element indicates the presence or absence of system dividers (also known as system separation marks) between systems displayed on the same page. Dividers on the left and right side of the page are controlled by the left-divider and right-divider elements respectively. The default vertical position is half the system-distance value from the top of the system that is below the divider. The default horizontal position is the left and right system margin, respectively.
//
// When used in the print element, the system-dividers element affects the dividers that would appear between the current system and the previous system.
type SystemDividers struct {
	XMLName      xml.Name                    `xml:"system-dividers"`
	LeftDivider  *EmptyPrintObjectStyleAlign `xml:"left-divider"`
	RightDivider *EmptyPrintObjectStyleAlign `xml:"right-divider"`
}

// SystemLayout is A system is a group of staves that are read and played simultaneously. System layout includes left and right margins and the vertical distance from the previous system. The system distance is measured from the bottom line of the previous system to the top line of the current system. It is ignored for the first system on a page. The top system distance is measured from the page's top margin to the top line of the first system. It is ignored for all but the first system on a page.
//
// Sometimes the sum of measure widths in a system may not equal the system width specified by the layout elements due to roundoff or other errors. The behavior when reading MusicXML files in these cases is application-dependent. For instance, applications may find that the system layout data is more reliable than the sum of the measure widths, and adjust the measure widths accordingly.
//
// When used in the defaults element, the system-layout element defines a default appearance for all systems in the score. If no system-layout element is present in the defaults element, default system layout values are chosen by the application.
//
// When used in the print element, the system-layout element affects the appearance of the current system only. All other systems use the default values as determined by the defaults element. If any child elements are missing from the system-layout element in a print element, the values determined by the defaults element are used there as well.
type SystemLayout struct {
	XMLName           xml.Name        `xml:"system-layout"`
	SystemMargins     *SystemMargins  `xml:"system-margins"`
	SystemDistance    float64         `xml:"system-distance"`
	TopSystemDistance float64         `xml:"top-system-distance"`
	SystemDividers    *SystemDividers `xml:"system-dividers"`
}

// SystemMargins is System margins are relative to the page margins. Positive values indent and negative values reduce the margin size.
type SystemMargins struct {
	XMLName          xml.Name `xml:"system-margins"`
	LeftRightMargins *LeftRightMargins
}

// Bookmark is The bookmark type serves as a well-defined target for an incoming simple XLink.
type Bookmark struct {
	XMLName         xml.Name `xml:"bookmark"`
	ElementPosition *ElementPosition
	IdAttr          string `xml:"id,attr"`
	NameAttr        string `xml:"name,attr,omitempty"`
}

// Link is The link type serves as an outgoing simple XLink. If a relative link is used within a document that is part of a compressed MusicXML file, the link is relative to the root folder of the zip file.
type Link struct {
	XMLName         xml.Name `xml:"link"`
	LinkAttributes  *LinkAttributes
	ElementPosition *ElementPosition
	Position        *Position
	NameAttr        string `xml:"name,attr,omitempty"`
}

// Accidental is The accidental type represents actual notated accidentals. Editorial and cautionary indications are indicated by attributes. Values for these attributes are "no" if not present. Specific graphic display such as parentheses, brackets, and size are controlled by the level-display attribute group.
type Accidental struct {
	XMLName        xml.Name `xml:"accidental"`
	LevelDisplay   *LevelDisplay
	PrintStyle     *PrintStyle
	CautionaryAttr string `xml:"cautionary,attr,omitempty"`
	EditorialAttr  string `xml:"editorial,attr,omitempty"`
	SmuflAttr      string `xml:"smufl,attr,omitempty"`
}

// AccidentalMark is An accidental-mark can be used as a separate notation or as part of an ornament. When used in an ornament, position and placement are relative to the ornament, not relative to the note.
type AccidentalMark struct {
	XMLName          xml.Name `xml:"accidental-mark"`
	LevelDisplay     *LevelDisplay
	PrintStyle       *PrintStyle
	Placement        *Placement
	OptionalUniqueId *OptionalUniqueId
	SmuflAttr        string `xml:"smufl,attr,omitempty"`
}

// Arpeggiate is The arpeggiate type indicates that this note is part of an arpeggiated chord. The number attribute can be used to distinguish between two simultaneous chords arpeggiated separately (different numbers) or together (same number). The up-down attribute is used if there is an arrow on the arpeggio sign. By default, arpeggios go from the lowest to highest note.  The length of the sign can be determined from the position attributes for the arpeggiate elements used with the top and bottom notes of the arpeggiated chord.
type Arpeggiate struct {
	XMLName          xml.Name `xml:"arpeggiate"`
	Position         *Position
	Placement        *Placement
	Color            string
	OptionalUniqueId *OptionalUniqueId
	NumberAttr       int    `xml:"number,attr,omitempty"`
	DirectionAttr    string `xml:"direction,attr,omitempty"`
}

// Articulations is The other-articulation element is used to define any articulations not yet in the MusicXML format. The smufl attribute can be used to specify a particular articulation, allowing application interoperability without requiring every SMuFL articulation to have a MusicXML element equivalent. Using the other-articulation element without the smufl attribute allows for extended representation, though without application interoperability.
type Articulations struct {
	XMLName           xml.Name `xml:"articulations"`
	OptionalUniqueId  *OptionalUniqueId
	Accent            *EmptyPlacement     `xml:"accent"`
	StrongAccent      *StrongAccent       `xml:"strong-accent"`
	Staccato          *EmptyPlacement     `xml:"staccato"`
	Tenuto            *EmptyPlacement     `xml:"tenuto"`
	DetachedLegato    *EmptyPlacement     `xml:"detached-legato"`
	Staccatissimo     *EmptyPlacement     `xml:"staccatissimo"`
	Spiccato          *EmptyPlacement     `xml:"spiccato"`
	Scoop             *EmptyLine          `xml:"scoop"`
	Plop              *EmptyLine          `xml:"plop"`
	Doit              *EmptyLine          `xml:"doit"`
	Falloff           *EmptyLine          `xml:"falloff"`
	BreathMark        *BreathMark         `xml:"breath-mark"`
	Caesura           *Caesura            `xml:"caesura"`
	Stress            *EmptyPlacement     `xml:"stress"`
	Unstress          *EmptyPlacement     `xml:"unstress"`
	SoftAccent        *EmptyPlacement     `xml:"soft-accent"`
	OtherArticulation *OtherPlacementText `xml:"other-articulation"`
}

// Arrow is The arrow element represents an arrow used for a musical technical indication. It can represent both Unicode and SMuFL arrows. The presence of an arrowhead element indicates that only the arrowhead is displayed, not the arrow stem. The smufl attribute distinguishes different SMuFL glyphs that have an arrow appearance such as arrowBlackUp, guitarStrumUp, or handbellsSwingUp. The specified glyph should match the descriptive representation.
type Arrow struct {
	XMLName        xml.Name `xml:"arrow"`
	PrintStyle     *PrintStyle
	Placement      *Placement
	Smufl          *Smufl
	ArrowDirection string `xml:"arrow-direction"`
	ArrowStyle     string `xml:"arrow-style"`
	Arrowhead      *Empty `xml:"arrowhead"`
	CircularArrow  string `xml:"circular-arrow"`
}

// Assess is By default, an assessment application should assess all notes without a cue child element, and not assess any note with a cue child element. The assess type allows this default assessment to be overridden for individual notes. The optional player and time-only attributes restrict the type to apply to a single player or set of times through a repeated section, respectively. If missing, the type applies to all players or all times through the repeated section, respectively. The player attribute references the id attribute of a player element defined within the matching score-part.
type Assess struct {
	XMLName      xml.Name `xml:"assess"`
	TypeAttr     string   `xml:"type,attr"`
	PlayerAttr   string   `xml:"player,attr,omitempty"`
	TimeOnlyAttr string   `xml:"time-only,attr,omitempty"`
}

// Backup is The backup and forward elements are required to coordinate multiple voices in one part, including music on multiple staves. The backup type is generally used to move between voices and staves. Thus the backup element does not include voice or staff elements. Duration values should always be positive, and should not cross measure boundaries or mid-measure changes in the divisions value.
type Backup struct {
	XMLName   xml.Name `xml:"backup"`
	Duration  string
	Editorial *Editorial
}

// Beam is Beam values include begin, continue, end, forward hook, and backward hook. Up to eight concurrent beams are available to cover up to 1024th notes. Each beam in a note is represented with a separate beam element, starting with the eighth note beam using a number attribute of 1.
//
// Note that the beam number does not distinguish sets of beams that overlap, as it does for slur and other elements. Beaming groups are distinguished by being in different voices and/or the presence or absence of grace and cue elements.
//
// Beams that have a begin value can also have a fan attribute to indicate accelerandos and ritardandos using fanned beams. The fan attribute may also be used with a continue value if the fanning direction changes on that note. The value is "none" if not specified.
//
// The repeater attribute has been deprecated in MusicXML 3.0. Formerly used for tremolos, it needs to be specified with a "yes" value for each beam using it.
type Beam struct {
	XMLName          xml.Name `xml:"beam"`
	Color            string
	OptionalUniqueId *OptionalUniqueId
	NumberAttr       int    `xml:"number,attr,omitempty"`
	RepeaterAttr     string `xml:"repeater,attr,omitempty"`
	FanAttr          string `xml:"fan,attr,omitempty"`
}

// Bend is The with-bar element indicates that the bend is to be done at the bridge with a whammy or vibrato bar. The content of the element indicates how this should be notated. Content values of "scoop" and "dip" refer to the SMuFL guitarVibratoBarScoop and guitarVibratoBarDip glyphs.
type Bend struct {
	XMLName    xml.Name `xml:"bend"`
	PrintStyle *PrintStyle
	BendSound  *BendSound
	ShapeAttr  string         `xml:"shape,attr,omitempty"`
	BendAlter  float64        `xml:"bend-alter"`
	PreBend    *Empty         `xml:"pre-bend"`
	Release    *Empty         `xml:"release"`
	WithBar    *PlacementText `xml:"with-bar"`
}

// BreathMark is The breath-mark element indicates a place to take a breath.
type BreathMark struct {
	XMLName    xml.Name `xml:"breath-mark"`
	PrintStyle *PrintStyle
	Placement  *Placement
}

// Caesura is The caesura element indicates a slight pause. It is notated using a "railroad tracks" symbol or other variations specified in the element content.
type Caesura struct {
	XMLName    xml.Name `xml:"caesura"`
	PrintStyle *PrintStyle
	Placement  *Placement
}

// Elision is The elision type represents an elision between lyric syllables. The text content specifies the symbol used to display the elision. Common values are a no-break space (Unicode 00A0), an underscore (Unicode 005F), or an undertie (Unicode 203F). If the text content is empty, the smufl attribute is used to specify the symbol to use. Its value is a SMuFL canonical glyph name that starts with lyrics. The SMuFL attribute is ignored if the elision glyph is already specified by the text content. If neither text content nor a smufl attribute are present, the elision glyph is application-specific.
type Elision struct {
	XMLName   xml.Name `xml:"elision"`
	Font      *Font
	Color     string
	SmuflAttr string `xml:"smufl,attr,omitempty"`
}

// EmptyLine is The empty-line type represents an empty element with line-shape, line-type, line-length, dashed-formatting, print-style and placement attributes.
type EmptyLine struct {
	XMLName          xml.Name `xml:"empty-line"`
	LineShape        string
	LineType         string
	LineLength       string
	DashedFormatting *DashedFormatting
	PrintStyle       *PrintStyle
	Placement        *Placement
}

// Extend is The extend type represents lyric word extension / melisma lines as well as figured bass extensions. The optional type and position attributes are added in Version 3.0 to provide better formatting control.
type Extend struct {
	XMLName  xml.Name `xml:"extend"`
	Position *Position
	Color    string
	TypeAttr string `xml:"type,attr,omitempty"`
}

// Figure is Values for the suffix element include plus and the accidental values sharp, flat, natural, double-sharp, flat-flat, and sharp-sharp. Suffixes include both symbols that come after the figure number and those that overstrike the figure number. The suffix values slash, back-slash, and vertical are used for slashed numbers indicating chromatic alteration. The orientation and display of the slash usually depends on the figure number. The suffix element may contain additional values for symbols specific to particular figured bass styles.
type Figure struct {
	XMLName      xml.Name `xml:"figure"`
	Editorial    *Editorial
	Prefix       *StyleText `xml:"prefix"`
	FigureNumber *StyleText `xml:"figure-number"`
	Suffix       *StyleText `xml:"suffix"`
	Extend       *Extend    `xml:"extend"`
}

// FiguredBass is The figured-bass element represents figured bass notation. Figured bass elements take their position from the first regular note (not a grace note or chord note) that follows in score order. The optional duration element is used to indicate changes of figures under a note.
//
// Figures are ordered from top to bottom. The value of parentheses is "no" if not present.
type FiguredBass struct {
	XMLName          xml.Name `xml:"figured-bass"`
	PrintStyleAlign  *PrintStyleAlign
	Placement        *Placement
	Printout         *Printout
	OptionalUniqueId *OptionalUniqueId
	ParenthesesAttr  string `xml:"parentheses,attr,omitempty"`
	Duration         string
	Editorial        *Editorial
	Figure           []*Figure `xml:"figure"`
}

// Forward is The backup and forward elements are required to coordinate multiple voices in one part, including music on multiple staves. The forward element is generally used within voices and staves. Duration values should always be positive, and should not cross measure boundaries or mid-measure changes in the divisions value.
type Forward struct {
	XMLName        xml.Name `xml:"forward"`
	Duration       string
	EditorialVoice *EditorialVoice
	Staff          *Staff
}

// Glissando is Glissando and slide types both indicate rapidly moving from one pitch to the other so that individual notes are not discerned. The distinction is similar to that between NIFF's glissando and portamento elements. A glissando sounds the half notes in between the slide and defaults to a wavy line. The optional text is printed alongside the line.
type Glissando struct {
	XMLName          xml.Name `xml:"glissando"`
	LineType         string
	DashedFormatting *DashedFormatting
	PrintStyle       *PrintStyle
	OptionalUniqueId *OptionalUniqueId
	TypeAttr         string `xml:"type,attr"`
	NumberAttr       int    `xml:"number,attr,omitempty"`
}

// Grace is The grace type indicates the presence of a grace note. The slash attribute for a grace note is yes for slashed eighth notes. The other grace note attributes come from MuseData sound suggestions. The steal-time-previous attribute indicates the percentage of time to steal from the previous note for the grace note. The steal-time-following attribute indicates the percentage of time to steal from the following note for the grace note, as for appoggiaturas. The make-time attribute indicates to make time, not steal time; the units are in real-time divisions for the grace note.
type Grace struct {
	XMLName                xml.Name `xml:"grace"`
	StealTimePreviousAttr  float64  `xml:"steal-time-previous,attr,omitempty"`
	StealTimeFollowingAttr float64  `xml:"steal-time-following,attr,omitempty"`
	MakeTimeAttr           float64  `xml:"make-time,attr,omitempty"`
	SlashAttr              string   `xml:"slash,attr,omitempty"`
}

// HammerOnPullOff is The hammer-on and pull-off elements are used in guitar and fretted instrument notation. Since a single slur can be marked over many notes, the hammer-on and pull-off elements are separate so the individual pair of notes can be specified. The element content can be used to specify how the hammer-on or pull-off should be notated. An empty element leaves this choice up to the application.
type HammerOnPullOff struct {
	XMLName    xml.Name `xml:"hammer-on-pull-off"`
	PrintStyle *PrintStyle
	Placement  *Placement
	TypeAttr   string `xml:"type,attr"`
	NumberAttr int    `xml:"number,attr,omitempty"`
}

// Handbell is The handbell element represents notation for various techniques used in handbell and handchime music.
type Handbell struct {
	XMLName    xml.Name `xml:"handbell"`
	PrintStyle *PrintStyle
	Placement  *Placement
}

// HarmonClosed is The harmon-closed type represents whether the harmon mute is closed, open, or half-open. The optional location attribute indicates which portion of the symbol is filled in when the element value is half.
type HarmonClosed struct {
	XMLName      xml.Name `xml:"harmon-closed"`
	LocationAttr string   `xml:"location,attr,omitempty"`
}

// HarmonMute is The harmon-mute type represents the symbols used for harmon mutes in brass notation.
type HarmonMute struct {
	XMLName      xml.Name `xml:"harmon-mute"`
	PrintStyle   *PrintStyle
	Placement    *Placement
	HarmonClosed *HarmonClosed `xml:"harmon-closed"`
}

// Harmonic is The sounding-pitch is the pitch which is heard when playing the harmonic.
type Harmonic struct {
	XMLName       xml.Name `xml:"harmonic"`
	PrintObject   *PrintObject
	PrintStyle    *PrintStyle
	Placement     *Placement
	Natural       *Empty `xml:"natural"`
	Artificial    *Empty `xml:"artificial"`
	BasePitch     *Empty `xml:"base-pitch"`
	TouchingPitch *Empty `xml:"touching-pitch"`
	SoundingPitch *Empty `xml:"sounding-pitch"`
}

// HeelToe is The heel and toe elements are used with organ pedals. The substitution value is "no" if the attribute is not present.
type HeelToe struct {
	XMLName          xml.Name `xml:"heel-toe"`
	SubstitutionAttr string   `xml:"substitution,attr,omitempty"`
}

// Hole is The optional hole-shape element indicates the shape of the hole symbol; the default is a circle.
type Hole struct {
	XMLName    xml.Name `xml:"hole"`
	PrintStyle *PrintStyle
	Placement  *Placement
	HoleType   string      `xml:"hole-type"`
	HoleClosed *HoleClosed `xml:"hole-closed"`
	HoleShape  string      `xml:"hole-shape"`
}

// HoleClosed is The hole-closed type represents whether the hole is closed, open, or half-open. The optional location attribute indicates which portion of the hole is filled in when the element value is half.
type HoleClosed struct {
	XMLName      xml.Name `xml:"hole-closed"`
	LocationAttr string   `xml:"location,attr,omitempty"`
}

// Instrument is The instrument type distinguishes between score-instrument elements in a score-part. The id attribute is an IDREF back to the score-instrument ID. If multiple score-instruments are specified on a score-part, there should be an instrument element for each note in the part.
type Instrument struct {
	XMLName xml.Name `xml:"instrument"`
	IdAttr  string   `xml:"id,attr"`
}

// Listen is The listen and listening types, new in Version 4.0, specify different ways that a score following or machine listening application can interact with a performer. The listen type handles interactions that are specific to a note. If multiple child elements of the same type are present, they should have distinct player and/or time-only attributes.
type Listen struct {
	XMLName     xml.Name        `xml:"listen"`
	Assess      *Assess         `xml:"assess"`
	Wait        *Wait           `xml:"wait"`
	OtherListen *OtherListening `xml:"other-listen"`
}

// Lyric is The end-paragraph element comes from RP-017 for Standard MIDI File Lyric meta-events. It facilitates lyric display for Karaoke and similar applications.
type Lyric struct {
	XMLName          xml.Name `xml:"lyric"`
	Justify          *Justify
	Position         *Position
	Placement        *Placement
	Color            string
	PrintObject      *PrintObject
	OptionalUniqueId *OptionalUniqueId
	NumberAttr       string `xml:"number,attr,omitempty"`
	NameAttr         string `xml:"name,attr,omitempty"`
	TimeOnlyAttr     string `xml:"time-only,attr,omitempty"`
	Editorial        *Editorial
	Syllabic         string           `xml:"syllabic"`
	Text             *TextElementData `xml:"text"`
	Elision          *Elision         `xml:"elision"`
	Extend           *Extend          `xml:"extend"`
	Laughing         *Empty           `xml:"laughing"`
	Humming          *Empty           `xml:"humming"`
	EndLine          *Empty           `xml:"end-line"`
	EndParagraph     *Empty           `xml:"end-paragraph"`
}

// Mordent is The mordent type is used for both represents the mordent sign with the vertical line and the inverted-mordent sign without the line. The long attribute is "no" by default. The approach and departure attributes are used for compound ornaments, indicating how the beginning and ending of the ornament look relative to the main part of the mordent.
type Mordent struct {
	XMLName       xml.Name `xml:"mordent"`
	LongAttr      string   `xml:"long,attr,omitempty"`
	ApproachAttr  string   `xml:"approach,attr,omitempty"`
	DepartureAttr string   `xml:"departure,attr,omitempty"`
}

// NonArpeggiate is The non-arpeggiate type indicates that this note is at the top or bottom of a bracket indicating to not arpeggiate these notes. Since this does not involve playback, it is only used on the top or bottom notes, not on each note as for the arpeggiate type.
type NonArpeggiate struct {
	XMLName          xml.Name `xml:"non-arpeggiate"`
	Position         *Position
	Placement        *Placement
	Color            string
	OptionalUniqueId *OptionalUniqueId
	TypeAttr         string `xml:"type,attr"`
	NumberAttr       int    `xml:"number,attr,omitempty"`
}

// Notations is Notations refer to musical notations, not XML notations. Multiple notations are allowed in order to represent multiple editorial levels. The print-object attribute, added in Version 3.0, allows notations to represent details of performance technique, such as fingerings, without having them appear in the score.
type Notations struct {
	XMLName          xml.Name `xml:"notations"`
	PrintObject      *PrintObject
	OptionalUniqueId *OptionalUniqueId
	Editorial        *Editorial
	Tied             *Tied           `xml:"tied"`
	Slur             *Slur           `xml:"slur"`
	Tuplet           *Tuplet         `xml:"tuplet"`
	Glissando        *Glissando      `xml:"glissando"`
	Slide            *Slide          `xml:"slide"`
	Ornaments        *Ornaments      `xml:"ornaments"`
	Technical        *Technical      `xml:"technical"`
	Articulations    *Articulations  `xml:"articulations"`
	Dynamics         *Dynamics       `xml:"dynamics"`
	Fermata          *Fermata        `xml:"fermata"`
	Arpeggiate       *Arpeggiate     `xml:"arpeggiate"`
	NonArpeggiate    *NonArpeggiate  `xml:"non-arpeggiate"`
	AccidentalMark   *AccidentalMark `xml:"accidental-mark"`
	OtherNotation    *OtherNotation  `xml:"other-notation"`
}

// Note is One dot element is used for each dot of prolongation. The placement element is used to specify whether the dot should appear above or below the staff line. It is ignored for notes that appear on a staff space.
type Note struct {
	XMLName          xml.Name `xml:"note"`
	XPosition        *XPosition
	Font             *Font
	Color            string
	Printout         *Printout
	OptionalUniqueId *OptionalUniqueId
	PrintLegerAttr   string  `xml:"print-leger,attr,omitempty"`
	DynamicsAttr     float64 `xml:"dynamics,attr,omitempty"`
	EndDynamicsAttr  float64 `xml:"end-dynamics,attr,omitempty"`
	AttackAttr       float64 `xml:"attack,attr,omitempty"`
	ReleaseAttr      float64 `xml:"release,attr,omitempty"`
	TimeOnlyAttr     string  `xml:"time-only,attr,omitempty"`
	PizzicatoAttr    string  `xml:"pizzicato,attr,omitempty"`
	FullNote         *FullNote
	Duration         string
	EditorialVoice   *EditorialVoice
	Staff            *Staff
	Grace            *Grace            `xml:"grace"`
	Tie              []*Tie            `xml:"tie"`
	Cue              *Empty            `xml:"cue"`
	Instrument       *Instrument       `xml:"instrument"`
	Type             *NoteType         `xml:"type"`
	Dot              []*EmptyPlacement `xml:"dot"`
	Accidental       *Accidental       `xml:"accidental"`
	TimeModification *TimeModification `xml:"time-modification"`
	Stem             *Stem             `xml:"stem"`
	Notehead         *Notehead         `xml:"notehead"`
	NoteheadText     *NoteheadText     `xml:"notehead-text"`
	Beam             []*Beam           `xml:"beam"`
	Notations        []*Notations      `xml:"notations"`
	Lyric            []*Lyric          `xml:"lyric"`
	Play             *Play             `xml:"play"`
	Listen           *Listen           `xml:"listen"`
}

// NoteType is The note-type type indicates the graphic note type. Values range from 1024th to maxima. The size attribute indicates full, cue, grace-cue, or large size. The default is full for regular notes, grace-cue for notes that contain both grace and cue elements, and cue for notes that contain either a cue or a grace element, but not both.
type NoteType struct {
	XMLName  xml.Name `xml:"note-type"`
	SizeAttr string   `xml:"size,attr,omitempty"`
}

// Notehead is The notehead type indicates shapes other than the open and closed ovals associated with note durations.
//
// The smufl attribute can be used to specify a particular notehead, allowing application interoperability without requiring every SMuFL glyph to have a MusicXML element equivalent. This attribute can be used either with the "other" value, or to refine a specific notehead value such as "cluster". Noteheads in the SMuFL "Note name noteheads" range (U+E150–U+E1AF) should not use the smufl attribute or the "other" value, but instead use the notehead-text element.
//
// For the enclosed shapes, the default is to be hollow for half notes and longer, and filled otherwise. The filled attribute can be set to change this if needed.
//
// If the parentheses attribute is set to yes, the notehead is parenthesized. It is no by default.
type Notehead struct {
	XMLName         xml.Name `xml:"notehead"`
	Font            *Font
	Color           string
	Smufl           *Smufl
	FilledAttr      string `xml:"filled,attr,omitempty"`
	ParenthesesAttr string `xml:"parentheses,attr,omitempty"`
}

// NoteheadText is The notehead-text type represents text that is displayed inside a notehead, as is done in some educational music. It is not needed for the numbers used in tablature or jianpu notation. The presence of a TAB or jianpu clefs is sufficient to indicate that numbers are used. The display-text and accidental-text elements allow display of fully formatted text and accidentals.
type NoteheadText struct {
	XMLName        xml.Name        `xml:"notehead-text"`
	DisplayText    *FormattedText  `xml:"display-text"`
	AccidentalText *AccidentalText `xml:"accidental-text"`
}

// Ornaments is The other-ornament element is used to define any ornaments not yet in the MusicXML format. The smufl attribute can be used to specify a particular ornament, allowing application interoperability without requiring every SMuFL ornament to have a MusicXML element equivalent. Using the other-ornament element without the smufl attribute allows for extended representation, though without application interoperability.
type Ornaments struct {
	XMLName              xml.Name `xml:"ornaments"`
	OptionalUniqueId     *OptionalUniqueId
	TrillMark            *EmptyTrillSound    `xml:"trill-mark"`
	Turn                 *HorizontalTurn     `xml:"turn"`
	DelayedTurn          *HorizontalTurn     `xml:"delayed-turn"`
	InvertedTurn         *HorizontalTurn     `xml:"inverted-turn"`
	DelayedInvertedTurn  *HorizontalTurn     `xml:"delayed-inverted-turn"`
	VerticalTurn         *EmptyTrillSound    `xml:"vertical-turn"`
	InvertedVerticalTurn *EmptyTrillSound    `xml:"inverted-vertical-turn"`
	Shake                *EmptyTrillSound    `xml:"shake"`
	WavyLine             *WavyLine           `xml:"wavy-line"`
	Mordent              *Mordent            `xml:"mordent"`
	InvertedMordent      *Mordent            `xml:"inverted-mordent"`
	Schleifer            *EmptyPlacement     `xml:"schleifer"`
	Tremolo              *Tremolo            `xml:"tremolo"`
	Haydn                *EmptyTrillSound    `xml:"haydn"`
	OtherOrnament        *OtherPlacementText `xml:"other-ornament"`
	AccidentalMark       []*AccidentalMark   `xml:"accidental-mark"`
}

// OtherNotation is The other-notation type is used to define any notations not yet in the MusicXML format. It handles notations where more specific extension elements such as other-dynamics and other-technical are not appropriate. The smufl attribute can be used to specify a particular notation, allowing application interoperability without requiring every SMuFL glyph to have a MusicXML element equivalent. Using the other-notation type without the smufl attribute allows for extended representation, though without application interoperability.
type OtherNotation struct {
	XMLName          xml.Name `xml:"other-notation"`
	PrintObject      *PrintObject
	PrintStyle       *PrintStyle
	Placement        *Placement
	Smufl            *Smufl
	OptionalUniqueId *OptionalUniqueId
	TypeAttr         string `xml:"type,attr"`
	NumberAttr       int    `xml:"number,attr,omitempty"`
}

// OtherPlacementText is The other-placement-text type represents a text element with print-style, placement, and smufl attribute groups. This type is used by MusicXML notation extension elements to allow specification of specific SMuFL glyphs without needed to add every glyph as a MusicXML element.
type OtherPlacementText struct {
	XMLName    xml.Name `xml:"other-placement-text"`
	PrintStyle *PrintStyle
	Placement  *Placement
	Smufl      *Smufl
}

// OtherText is The other-text type represents a text element with a smufl attribute group. This type is used by MusicXML direction extension elements to allow specification of specific SMuFL glyphs without needed to add every glyph as a MusicXML element.
type OtherText struct {
	XMLName xml.Name `xml:"other-text"`
	Smufl   *Smufl
}

// Pitch is Pitch is represented as a combination of the step of the diatonic scale, the chromatic alteration, and the octave.
type Pitch struct {
	XMLName xml.Name `xml:"pitch"`
	Step    string   `xml:"step"`
	Alter   float64  `xml:"alter"`
	Octave  int      `xml:"octave"`
}

// PlacementText is The placement-text type represents a text element with print-style and placement attribute groups.
type PlacementText struct {
	XMLName    xml.Name `xml:"placement-text"`
	PrintStyle *PrintStyle
	Placement  *Placement
}

// Rest is The rest element indicates notated rests or silences. Rest elements are usually empty, but placement on the staff can be specified using display-step and display-octave elements. If the measure attribute is set to yes, this indicates this is a complete measure rest.
type Rest struct {
	XMLName           xml.Name `xml:"rest"`
	MeasureAttr       string   `xml:"measure,attr,omitempty"`
	DisplayStepOctave *DisplayStepOctave
}

// Slide is Glissando and slide types both indicate rapidly moving from one pitch to the other so that individual notes are not discerned. The distinction is similar to that between NIFF's glissando and portamento elements. A slide is continuous between two notes and defaults to a solid line. The optional text for a is printed alongside the line.
type Slide struct {
	XMLName          xml.Name `xml:"slide"`
	LineType         string
	DashedFormatting *DashedFormatting
	PrintStyle       *PrintStyle
	BendSound        *BendSound
	OptionalUniqueId *OptionalUniqueId
	TypeAttr         string `xml:"type,attr"`
	NumberAttr       int    `xml:"number,attr,omitempty"`
}

// Slur is Slur types are empty. Most slurs are represented with two elements: one with a start type, and one with a stop type. Slurs can add more elements using a continue type. This is typically used to specify the formatting of cross-system slurs, or to specify the shape of very complex slurs.
type Slur struct {
	XMLName          xml.Name `xml:"slur"`
	LineType         string
	DashedFormatting *DashedFormatting
	Position         *Position
	Placement        *Placement
	Orientation      *Orientation
	Bezier           *Bezier
	Color            string
	OptionalUniqueId *OptionalUniqueId
	TypeAttr         string `xml:"type,attr"`
	NumberAttr       int    `xml:"number,attr,omitempty"`
}

// Stem is Stems can be down, up, none, or double. For down and up stems, the position attributes can be used to specify stem length. The relative values specify the end of the stem relative to the program default. Default values specify an absolute end stem position. Negative values of relative-y that would flip a stem instead of shortening it are ignored. A stem element associated with a rest refers to a stemlet.
type Stem struct {
	XMLName   xml.Name `xml:"stem"`
	YPosition *YPosition
	Color     string
}

// StrongAccent is The strong-accent type indicates a vertical accent mark. The type attribute indicates if the point of the accent is down or up.
type StrongAccent struct {
	XMLName  xml.Name `xml:"strong-accent"`
	TypeAttr string   `xml:"type,attr,omitempty"`
}

// StyleText is The style-text type represents a text element with a print-style attribute group.
type StyleText struct {
	XMLName    xml.Name `xml:"style-text"`
	PrintStyle *PrintStyle
}

// Tap is The tap type indicates a tap on the fretboard. The text content allows specification of the notation; + and T are common choices. If the element is empty, the hand attribute is used to specify the symbol to use. The hand attribute is ignored if the tap glyph is already specified by the text content. If neither text content nor the hand attribute are present, the display is application-specific.
type Tap struct {
	XMLName    xml.Name `xml:"tap"`
	PrintStyle *PrintStyle
	Placement  *Placement
	HandAttr   string `xml:"hand,attr,omitempty"`
}

// Technical is The other-technical element is used to define any technical indications not yet in the MusicXML format. The smufl attribute can be used to specify a particular glyph, allowing application interoperability without requiring every SMuFL technical indication to have a MusicXML element equivalent. Using the other-technical element without the smufl attribute allows for extended representation, though without application interoperability.
type Technical struct {
	XMLName          xml.Name `xml:"technical"`
	OptionalUniqueId *OptionalUniqueId
	UpBow            *EmptyPlacement      `xml:"up-bow"`
	DownBow          *EmptyPlacement      `xml:"down-bow"`
	Harmonic         *Harmonic            `xml:"harmonic"`
	OpenString       *EmptyPlacement      `xml:"open-string"`
	ThumbPosition    *EmptyPlacement      `xml:"thumb-position"`
	Fingering        *Fingering           `xml:"fingering"`
	Pluck            *PlacementText       `xml:"pluck"`
	DoubleTongue     *EmptyPlacement      `xml:"double-tongue"`
	TripleTongue     *EmptyPlacement      `xml:"triple-tongue"`
	Stopped          *EmptyPlacementSmufl `xml:"stopped"`
	SnapPizzicato    *EmptyPlacement      `xml:"snap-pizzicato"`
	Fret             *Fret                `xml:"fret"`
	String           string               `xml:"string"`
	HammerOn         *HammerOnPullOff     `xml:"hammer-on"`
	PullOff          *HammerOnPullOff     `xml:"pull-off"`
	Bend             *Bend                `xml:"bend"`
	Tap              *Tap                 `xml:"tap"`
	Heel             *HeelToe             `xml:"heel"`
	Toe              *HeelToe             `xml:"toe"`
	Fingernails      *EmptyPlacement      `xml:"fingernails"`
	Hole             *Hole                `xml:"hole"`
	Arrow            *Arrow               `xml:"arrow"`
	Handbell         *Handbell            `xml:"handbell"`
	BrassBend        *EmptyPlacement      `xml:"brass-bend"`
	Flip             *EmptyPlacement      `xml:"flip"`
	Smear            *EmptyPlacement      `xml:"smear"`
	Open             *EmptyPlacementSmufl `xml:"open"`
	HalfMuted        *EmptyPlacementSmufl `xml:"half-muted"`
	HarmonMute       *HarmonMute          `xml:"harmon-mute"`
	Golpe            *EmptyPlacement      `xml:"golpe"`
	OtherTechnical   *OtherPlacementText  `xml:"other-technical"`
}

// TextElementData is The text-element-data type represents a syllable or portion of a syllable for lyric text underlay. A hyphen in the string content should only be used for an actual hyphenated word. Language names for text elements come from ISO 639, with optional country subcodes from ISO 3166.
type TextElementData struct {
	XMLName        xml.Name `xml:"text-element-data"`
	Font           *Font
	Color          string
	TextDecoration *TextDecoration
	TextRotation   *TextRotation
	LetterSpacing  *LetterSpacing
	TextDirection  string
	XmlLangAttr    *Lang `xml:"xml:lang,attr,omitempty"`
}

// Tie is The tie element indicates that a tie begins or ends with this note. If the tie element applies only particular times through a repeat, the time-only attribute indicates which times to apply it. The tie element indicates sound; the tied element indicates notation.
type Tie struct {
	XMLName      xml.Name `xml:"tie"`
	TypeAttr     string   `xml:"type,attr"`
	TimeOnlyAttr string   `xml:"time-only,attr,omitempty"`
}

// Tied is The tied element represents the notated tie. The tie element represents the tie sound.
//
// The number attribute is rarely needed to disambiguate ties, since note pitches will usually suffice. The attribute is implied rather than defaulting to 1 as with most elements. It is available for use in more complex tied notation situations.
//
// Ties that join two notes of the same pitch together should be represented with a tied element on the first note with type="start" and a tied element on the second note with type="stop".  This can also be done if the two notes being tied are enharmonically equivalent, but have different step values. It is not recommended to use tied elements to join two notes with enharmonically inequivalent pitches.
//
// Ties that indicate that an instrument should be undamped are specified with a single tied element with type="let-ring".
//
// Ties that are visually attached to only one note, other than undamped ties, should be specified with two tied elements on the same note, first type="start" then type="stop". This can be used to represent ties into or out of repeated sections or codas.
type Tied struct {
	XMLName          xml.Name `xml:"tied"`
	LineType         string
	DashedFormatting *DashedFormatting
	Position         *Position
	Placement        *Placement
	Orientation      *Orientation
	Bezier           *Bezier
	Color            string
	OptionalUniqueId *OptionalUniqueId
	TypeAttr         string `xml:"type,attr"`
	NumberAttr       int    `xml:"number,attr,omitempty"`
}

// TimeModification is The normal-dot element is used to specify dotted normal tuplet types.
type TimeModification struct {
	XMLName     xml.Name `xml:"time-modification"`
	ActualNotes int      `xml:"actual-notes"`
	NormalNotes int      `xml:"normal-notes"`
	NormalType  string   `xml:"normal-type"`
	NormalDot   []*Empty `xml:"normal-dot"`
}

// Tremolo is The tremolo ornament can be used to indicate single-note, double-note, or unmeasured tremolos. Single-note tremolos use the single type, double-note tremolos use the start and stop types, and unmeasured tremolos use the unmeasured type. The default is "single" for compatibility with Version 1.1. The text of the element indicates the number of tremolo marks and is an integer from 0 to 8. Note that the number of attached beams is not included in this value, but is represented separately using the beam element. The value should be 0 for unmeasured tremolos.
//
// When using double-note tremolos, the duration of each note in the tremolo should correspond to half of the notated type value. A time-modification element should also be added with an actual-notes value of 2 and a normal-notes value of 1. If used within a tuplet, this 2/1 ratio should be multiplied by the existing tuplet ratio.
//
// The smufl attribute specifies the glyph to use from the SMuFL tremolos range for an unmeasured tremolo. It is ignored for other tremolo types. The SMuFL buzzRoll glyph is used by default if the attribute is missing.
//
// Using repeater beams for indicating tremolos is deprecated as of MusicXML 3.0.
type Tremolo struct {
	XMLName    xml.Name `xml:"tremolo"`
	PrintStyle *PrintStyle
	Placement  *Placement
	Smufl      *Smufl
	TypeAttr   string `xml:"type,attr,omitempty"`
}

// Tuplet is The tuplet-normal element provide optional full control over how the normal part of the tuplet is displayed, including number and note type (with dots). If any of these elements are absent, their values are based on the time-modification element.
type Tuplet struct {
	XMLName          xml.Name `xml:"tuplet"`
	LineShape        string
	Position         *Position
	Placement        *Placement
	OptionalUniqueId *OptionalUniqueId
	TypeAttr         string         `xml:"type,attr"`
	NumberAttr       int            `xml:"number,attr,omitempty"`
	BracketAttr      string         `xml:"bracket,attr,omitempty"`
	ShowNumberAttr   string         `xml:"show-number,attr,omitempty"`
	ShowTypeAttr     string         `xml:"show-type,attr,omitempty"`
	TupletActual     *TupletPortion `xml:"tuplet-actual"`
	TupletNormal     *TupletPortion `xml:"tuplet-normal"`
}

// TupletDot is The tuplet-dot type is used to specify dotted normal tuplet types.
type TupletDot struct {
	XMLName xml.Name `xml:"tuplet-dot"`
	Font    *Font
	Color   string
}

// TupletNumber is The tuplet-number type indicates the number of notes for this portion of the tuplet.
type TupletNumber struct {
	XMLName xml.Name `xml:"tuplet-number"`
	Font    *Font
	Color   string
}

// TupletPortion is The tuplet-portion type provides optional full control over tuplet specifications. It allows the number and note type (including dots) to be set for the actual and normal portions of a single tuplet. If any of these elements are absent, their values are based on the time-modification element.
type TupletPortion struct {
	XMLName      xml.Name      `xml:"tuplet-portion"`
	TupletNumber *TupletNumber `xml:"tuplet-number"`
	TupletType   *TupletType   `xml:"tuplet-type"`
	TupletDot    []*TupletDot  `xml:"tuplet-dot"`
}

// TupletType is The tuplet-type type indicates the graphical note type of the notes for this portion of the tuplet.
type TupletType struct {
	XMLName xml.Name `xml:"tuplet-type"`
	Font    *Font
	Color   string
}

// Unpitched is The unpitched type represents musical elements that are notated on the staff but lack definite pitch, such as unpitched percussion and speaking voice.
type Unpitched struct {
	XMLName           xml.Name `xml:"unpitched"`
	DisplayStepOctave *DisplayStepOctave
}

// Wait is The wait type specifies a point where the accompaniment should wait for a performer event before continuing. This typically happens at the start of new sections or after a held note or indeterminate music. These waiting points cannot always be inferred reliably from the contents of the displayed score. The optional player and time-only attributes restrict the type to apply to a single player or set of times through a repeated section, respectively.
type Wait struct {
	XMLName      xml.Name `xml:"wait"`
	PlayerAttr   string   `xml:"player,attr,omitempty"`
	TimeOnlyAttr string   `xml:"time-only,attr,omitempty"`
}

// Credit is The credit type represents the appearance of the title, composer, arranger, lyricist, copyright, dedication, and other text, symbols, and graphics that commonly appear on the first page of a score. The credit-words, credit-symbol, and credit-image elements are similar to the words, symbol, and image elements for directions. However, since the credit is not part of a measure, the default-x and default-y attributes adjust the origin relative to the bottom left-hand corner of the page. The enclosure for credit-words and credit-symbol is none by default.
//
// By default, a series of credit-words and credit-symbol elements within a single credit element follow one another in sequence visually. Non-positional formatting attributes are carried over from the previous element by default.
//
// The page attribute for the credit element specifies the page number where the credit should appear. This is an integer value that starts with 1 for the first page. Its value is 1 by default. Since credits occur before the music, these page numbers do not refer to the page numbering specified by the print element's page-number attribute.
//
// The credit-type element indicates the purpose behind a credit. Multiple types of data may be combined in a single credit, so multiple elements may be used. Standard values include page number, title, subtitle, composer, arranger, lyricist, rights, and part name.
type Credit struct {
	XMLName          xml.Name `xml:"credit"`
	OptionalUniqueId *OptionalUniqueId
	PageAttr         int                `xml:"page,attr,omitempty"`
	CreditType       []string           `xml:"credit-type"`
	Link             []*Link            `xml:"link"`
	Bookmark         []*Bookmark        `xml:"bookmark"`
	CreditImage      *Image             `xml:"credit-image"`
	CreditWords      *FormattedTextId   `xml:"credit-words"`
	CreditSymbol     *FormattedSymbolId `xml:"credit-symbol"`
}

// Defaults is The presence of a concert score element indicates that a score is displayed in concert pitch. It is used for scores that contain parts for transposing instruments.
//
// A file with a concert-score element may not contain any transpose elements that have non-zero values for either the diatonic or chromatic elements. Concert scores may include octave transpositions, so transpose elements with a double element or a non-zero octave-change element value are permitted.
type Defaults struct {
	XMLName       xml.Name `xml:"defaults"`
	Layout        *Layout
	Scaling       *Scaling         `xml:"scaling"`
	ConcertScore  *Empty           `xml:"concert-score"`
	Appearance    *Appearance      `xml:"appearance"`
	MusicFont     *EmptyFont       `xml:"music-font"`
	WordFont      *EmptyFont       `xml:"word-font"`
	LyricFont     []*LyricFont     `xml:"lyric-font"`
	LyricLanguage []*LyricLanguage `xml:"lyric-language"`
}

// EmptyFont is The empty-font type represents an empty element with font attributes.
type EmptyFont struct {
	XMLName xml.Name `xml:"empty-font"`
	Font    *Font
}

// GroupBarline is The group-barline type indicates if the group should have common barlines.
type GroupBarline struct {
	XMLName xml.Name `xml:"group-barline"`
	Color   string
}

// GroupName is The group-name type describes the name or abbreviation of a part-group element. Formatting attributes in the group-name type are deprecated in Version 2.0 in favor of the new group-name-display and group-abbreviation-display elements.
type GroupName struct {
	XMLName       xml.Name `xml:"group-name"`
	GroupNameText *GroupNameText
}

// GroupSymbol is The group-symbol type indicates how the symbol for a group is indicated in the score.
type GroupSymbol struct {
	XMLName  xml.Name `xml:"group-symbol"`
	Position *Position
	Color    string
}

// LyricFont is The lyric-font type specifies the default font for a particular name and number of lyric.
type LyricFont struct {
	XMLName    xml.Name `xml:"lyric-font"`
	Font       *Font
	NumberAttr string `xml:"number,attr,omitempty"`
	NameAttr   string `xml:"name,attr,omitempty"`
}

// LyricLanguage is The lyric-language type specifies the default language for a particular name and number of lyric.
type LyricLanguage struct {
	XMLName     xml.Name `xml:"lyric-language"`
	NumberAttr  string   `xml:"number,attr,omitempty"`
	NameAttr    string   `xml:"name,attr,omitempty"`
	XmlLangAttr *Lang    `xml:"xml:lang,attr"`
}

// Opus is The opus type represents a link to a MusicXML opus document that composes multiple MusicXML scores into a collection.
type Opus struct {
	XMLName        xml.Name `xml:"opus"`
	LinkAttributes *LinkAttributes
}

// PartGroup is The group-time element indicates that the displayed time signatures should stretch across all parts and staves in the group.
type PartGroup struct {
	XMLName                  xml.Name `xml:"part-group"`
	TypeAttr                 string   `xml:"type,attr"`
	NumberAttr               string   `xml:"number,attr,omitempty"`
	Editorial                *Editorial
	GroupName                *GroupName    `xml:"group-name"`
	GroupNameDisplay         *NameDisplay  `xml:"group-name-display"`
	GroupAbbreviation        *GroupName    `xml:"group-abbreviation"`
	GroupAbbreviationDisplay *NameDisplay  `xml:"group-abbreviation-display"`
	GroupSymbol              *GroupSymbol  `xml:"group-symbol"`
	GroupBarline             *GroupBarline `xml:"group-barline"`
	GroupTime                *Empty        `xml:"group-time"`
}

// PartLink is The part-link type allows MusicXML data for both score and parts to be contained within a single compressed MusicXML file. It links a score-part from a score file to a MusicXML files that contain parts data. In the case of a single compressed MusicXML file, the link href values are paths that are relative to the root folder of the zip file.
//
// The score-part containing the part-link should include a group element with a score value. The MusicXML file that is the target of the part-link should have score-part elements that include a group element with a parts value.
//
// Multiple part-link elements can link a condensed part within a score file to multiple MusicXML parts files. For example, a "Clarinet 1 and 2" part in a score file could link to separate "Clarinet 1" and "Clarinet 2" part files.
type PartLink struct {
	XMLName        xml.Name `xml:"part-link"`
	LinkAttributes *LinkAttributes
}

// PartList is The part-list identifies the different musical parts in this movement. Each part has an ID that is used later within the musical data. Since parts may be encoded separately and combined later, identification elements are present at both the score and score-part levels. There must be at least one score-part, combined as desired with part-group elements that indicate braces and brackets. Parts are ordered from top to bottom in a score based on the order in which they appear in the part-list.
type PartList struct {
	XMLName   xml.Name `xml:"part-list"`
	PartGroup []*PartGroup
	ScorePart *ScorePart
}

// PartName is The part-name type describes the name or abbreviation of a score-part element. Formatting attributes for the part-name element are deprecated in Version 2.0 in favor of the new part-name-display and part-abbreviation-display elements.
type PartName struct {
	XMLName      xml.Name `xml:"part-name"`
	PartNameText *PartNameText
}

// Player is The player-name element is typically used within a software application, rather than appearing on the printed page of a score.
type Player struct {
	XMLName    xml.Name `xml:"player"`
	IdAttr     string   `xml:"id,attr"`
	PlayerName string   `xml:"player-name"`
}

// ScoreInstrument is The optional instrument-abbreviation element is typically used within a software application, rather than appearing on the printed page of a score.
type ScoreInstrument struct {
	XMLName                xml.Name `xml:"score-instrument"`
	IdAttr                 string   `xml:"id,attr"`
	VirtualInstrumentData  *VirtualInstrumentData
	InstrumentName         string `xml:"instrument-name"`
	InstrumentAbbreviation string `xml:"instrument-abbreviation"`
}

// ScorePart is The group element allows the use of different versions of the part for different purposes. Typical values include score, parts, sound, and data. Ordering information that is directly encoded in MuseData can be derived from the ordering within a MusicXML score or opus.
type ScorePart struct {
	XMLName                 xml.Name           `xml:"score-part"`
	IdAttr                  string             `xml:"id,attr"`
	Identification          *Identification    `xml:"identification"`
	PartLink                []*PartLink        `xml:"part-link"`
	PartName                *PartName          `xml:"part-name"`
	PartNameDisplay         *NameDisplay       `xml:"part-name-display"`
	PartAbbreviation        *PartName          `xml:"part-abbreviation"`
	PartAbbreviationDisplay *NameDisplay       `xml:"part-abbreviation-display"`
	Group                   []string           `xml:"group"`
	ScoreInstrument         []*ScoreInstrument `xml:"score-instrument"`
	Player                  []*Player          `xml:"player"`
	MidiDevice              *MidiDevice        `xml:"midi-device"`
	MidiInstrument          *MidiInstrument    `xml:"midi-instrument"`
}

// VirtualInstrument is The virtual-name element indicates the library-specific name for the virtual instrument.
type VirtualInstrument struct {
	XMLName        xml.Name `xml:"virtual-instrument"`
	VirtualLibrary string   `xml:"virtual-library"`
	VirtualName    string   `xml:"virtual-name"`
}

// Work is The work-title element specifies the title of a work, not including its opus or other work number.
type Work struct {
	XMLName    xml.Name `xml:"work"`
	WorkNumber string   `xml:"work-number"`
	WorkTitle  string   `xml:"work-title"`
	Opus       *Opus    `xml:"opus"`
}

// Editorial ...
type Editorial struct {
	XMLName  xml.Name `xml:"editorial"`
	Footnote *Footnote
	Level    *Level
}

// EditorialVoice ...
type EditorialVoice struct {
	XMLName  xml.Name `xml:"editorial-voice"`
	Footnote *Footnote
	Level    *Level
	Voice    *Voice
}

// EditorialVoiceDirection ...
type EditorialVoiceDirection struct {
	XMLName  xml.Name `xml:"editorial-voice-direction"`
	Footnote *Footnote
	Level    *Level
	Voice    *Voice
}

// Footnote ...
type Footnote struct {
	XMLName  xml.Name `xml:"footnote"`
	Footnote *FormattedText
}

// Staff ...
type Staff struct {
	XMLName xml.Name `xml:"staff"`
	Staff   int
}

// Tuning ...
type Tuning struct {
	XMLName      xml.Name `xml:"tuning"`
	TuningStep   string
	TuningAlter  float64
	TuningOctave int
}

// VirtualInstrumentData ...
type VirtualInstrumentData struct {
	XMLName           xml.Name `xml:"virtual-instrument-data"`
	InstrumentSound   string
	Solo              *Empty
	Ensemble          *PositiveIntegerOrEmpty
	VirtualInstrument *VirtualInstrument
}

// Voice ...
type Voice struct {
	XMLName xml.Name `xml:"voice"`
	Voice   string
}

// NonTraditionalKey ...
type NonTraditionalKey struct {
	XMLName       xml.Name `xml:"non-traditional-key"`
	KeyStep       string
	KeyAlter      float64
	KeyAccidental *KeyAccidental
}

// TimeSignature ...
type TimeSignature struct {
	XMLName  xml.Name `xml:"time-signature"`
	Beats    string
	BeatType string
}

// TraditionalKey ...
type TraditionalKey struct {
	XMLName xml.Name `xml:"traditional-key"`
	Cancel  *Cancel
	Fifths  int
	Mode    string
}

// BeatUnit ...
type BeatUnit struct {
	XMLName     xml.Name `xml:"beat-unit"`
	BeatUnit    string
	BeatUnitDot []*Empty
}

// HarmonyChord ...
type HarmonyChord struct {
	XMLName   xml.Name `xml:"harmony-chord"`
	Root      *Root
	Function  *StyleText
	Kind      *Kind
	Inversion *Inversion
	Bass      *Bass
	Degree    []*Degree
}

// AllMargins ...
type AllMargins struct {
	XMLName          xml.Name `xml:"all-margins"`
	TopMargin        float64
	BottomMargin     float64
	LeftRightMargins *LeftRightMargins
}

// Layout ...
type Layout struct {
	XMLName      xml.Name `xml:"layout"`
	PageLayout   *PageLayout
	SystemLayout *SystemLayout
	StaffLayout  []*StaffLayout
}

// LeftRightMargins ...
type LeftRightMargins struct {
	XMLName     xml.Name `xml:"left-right-margins"`
	LeftMargin  float64
	RightMargin float64
}

// Duration ...
type Duration struct {
	XMLName  xml.Name `xml:"duration"`
	Duration float64
}

// DisplayStepOctave ...
type DisplayStepOctave struct {
	XMLName       xml.Name `xml:"display-step-octave"`
	DisplayStep   string
	DisplayOctave int
}

// FullNote ...
type FullNote struct {
	XMLName   xml.Name `xml:"full-note"`
	Chord     *Empty
	Pitch     *Pitch
	Unpitched *Unpitched
	Rest      *Rest
}

// MusicData ...
type MusicData struct {
	XMLName     xml.Name `xml:"music-data"`
	Note        *Note
	Backup      *Backup
	Forward     *Forward
	Direction   *Direction
	Attributes  *Attributes
	Harmony     *Harmony
	FiguredBass *FiguredBass
	Print       *Print
	Sound       *Sound
	Listening   *Listening
	Barline     *Barline
	Grouping    *Grouping
	Link        *Link
	Bookmark    *Bookmark
}

// ScoreHeader ...
type ScoreHeader struct {
	XMLName        xml.Name `xml:"score-header"`
	Work           *Work
	MovementNumber string
	MovementTitle  string
	Identification *Identification
	Defaults       *Defaults
	Credit         []*Credit
	PartList       *PartList
}

// Measure ...
type Measure struct {
	XMLName           xml.Name `xml:"measure"`
	MeasureAttributes *MeasureAttributes
	MusicData         *MusicData
}

// Part ...
type Part struct {
	XMLName        xml.Name `xml:"part"`
	PartAttributes *PartAttributes
	Measure        []*Measure `xml:"measure"`
}

// ScorePartwise ...
type ScorePartwise struct {
	XMLName            xml.Name `xml:"score-partwise"`
	DocumentAttributes *DocumentAttributes
	ScoreHeader        *ScoreHeader
	Part               []*Part `xml:"part"`
}

// ScoreTimewise ...
type ScoreTimewise struct {
	XMLName            xml.Name `xml:"score-timewise"`
	DocumentAttributes *DocumentAttributes
	ScoreHeader        *ScoreHeader
	Measure            []*Measure `xml:"measure"`
}
